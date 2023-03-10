// Code generated by hertz generator.

package core

import (
	"common/utils"
	"context"

	"common/kitex_gen/douyin/rpc"
	core "gateway/biz/model/douyin/core"
	"gateway/internal/cache"
	"gateway/internal/db"
	"gateway/internal/jwt"
	serivces "gateway/internal/services"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func toUserList(ctx context.Context, ids []int64, following int64) ([]*core.User, error) {
	users, err := serivces.GatherUserInfoFromIds(ctx, following, ids, nil, true, following != 0)
	if err != nil {
		return nil, err
	}
	converted := make([]*core.User, 0, len(users))
	for _, id := range ids {
		user := users[id]
		if following == 0 {
			user.IsFollow = true
		}
		converted = append(converted, user)
	}
	return converted, nil
}

// Relation .
// @router /douyin/relation/action/ [POST]
func Relation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	user, _, err := jwt.Validate(req.Token)
	if err != nil && user != 0 {
		utils.ErrorUnauthorized.Write(c)
		return
	}

	if !db.UserExists(req.ToUserId) {
		utils.ErrorNoSuchUser.Write(c)
		return
	}

	r, err := serivces.Feed.Relation(ctx, &rpc.DouyinRelationActionRequest{
		RequestUserId: user,
		ToUserId:      req.ToUserId,
		ActionType:    req.ActionType,
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	resp := new(core.DouyinRelationActionResponse)

	cache.Flush(user)
	cache.Flush(req.ToUserId)
	cache.FlushFollowing(user)

	c.JSON(consts.StatusOK, resp)
}

// Following .
// @router /douyin/relation/follow/list/ [GET]
func Following(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	user, err := jwt.AuthorizedUser(c, &req.Token)
	if err != nil {
		utils.Error(c, err)
		return
	}

	if !db.UserExists(req.UserId) {
		utils.ErrorNoSuchUser.Write(c)
		return
	}

	var ids []int64
	if ids = cache.GetFollowing(user); ids == nil {
		r, err := serivces.Feed.Following(ctx, &rpc.DouyinRelationFollowListRequest{
			UserId:        req.UserId,
			RequestUserId: user,
		})
		if err != nil {
			hlog.Warn(err)
			utils.ErrorRpcTimeout.Write(c)
			return
		}
		if utils.RpcError(c, r.StatusCode) {
			return
		}
		ids = r.UserList
		cache.SetFollowing(user, ids)
	}

	list, err := toUserList(ctx, ids, 0)
	if err != nil {
		utils.Error(c, err)
		return
	}
	resp := &core.DouyinRelationFollowListResponse{
		UserList: list,
	}

	c.JSON(consts.StatusOK, resp)
}

// Follower .
// @router /douyin/relation/follower/list/ [GET]
func Follower(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	user, err := jwt.AuthorizedUser(c, &req.Token)
	if err != nil {
		utils.Error(c, err)
		return
	}

	if !db.UserExists(req.UserId) {
		utils.ErrorNoSuchUser.Write(c)
		return
	}

	r, err := serivces.Feed.Follower(ctx, &rpc.DouyinRelationFollowerListRequest{
		UserId:        req.UserId,
		RequestUserId: user,
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	list, err := toUserList(ctx, r.UserList, user)
	if err != nil {
		utils.Error(c, err)
		return
	}
	resp := &core.DouyinRelationFollowerListResponse{
		UserList: list,
	}

	c.JSON(consts.StatusOK, resp)
}

// Friend .
// @router /douyin/relation/friend/list/ [GET]
func Friend(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	user, err := jwt.AuthorizedUser(c, &req.Token)
	if err != nil {
		utils.Error(c, err)
		return
	}

	if !db.UserExists(req.UserId) {
		utils.ErrorNoSuchUser.Write(c)
		return
	}

	r, err := serivces.Feed.Friend(ctx, &rpc.DouyinRelationFriendListRequest{
		UserId:        req.UserId,
		RequestUserId: user,
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	list, err := toUserList(ctx, r.UserList, 0)
	if err != nil {
		utils.Error(c, err)
		return
	}

	if err := fillMessages(ctx, user, r.UserList, list); err != utils.ErrorOk {
		err.Write(c)
		return
	}

	resp := &core.DouyinRelationFriendListResponse{
		UserList: list,
	}

	c.JSON(consts.StatusOK, resp)
}

func fillMessages(ctx context.Context, user int64, friends []int64, users []*core.User) utils.ErrorCode {
	r, err := serivces.Message.LatestMessages(ctx, &rpc.LatestMessageRequest{
		Friends:       friends,
		RequestUserId: user,
	})
	if err != nil {
		return utils.ErrorRpcTimeout
	}
	mapping := make(map[int64]int)
	for i, m := range r.Messages {
		id := m.FromUserId
		if id == user {
			id = m.ToUserId
		}
		mapping[id] = i
	}

	for _, u := range users {
		if i, ok := mapping[u.Id]; ok {
			var msgType int64
			if r.Messages[i].FromUserId == user {
				msgType = 1
			}
			u.MsgType = &msgType
			u.Message = &r.Messages[i].Content
		}
	}
	return utils.ErrorOk
}
