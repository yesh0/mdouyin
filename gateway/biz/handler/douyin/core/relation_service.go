// Code generated by hertz generator.

package core

import (
	"common/utils"
	"context"

	"common/kitex_gen/douyin/rpc"
	core "gateway/biz/model/douyin/core"
	"gateway/internal/db"
	"gateway/internal/jwt"
	serivces "gateway/internal/services"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func toUserList(ids []int64) ([]*core.User, error) {
	users, err := db.FindUsersByIds(ids)
	if err != nil {
		return nil, utils.ErrorInternalError
	}
	converted := make([]*core.User, 0, len(users))
	for _, u := range users {
		converted = append(converted, &core.User{
			Id:   int64(u.Id),
			Name: u.Name,
		})
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
	if err != nil {
		utils.ErrorUnauthorized.Write(c)
		return
	}

	if !db.UserExists(req.ToUserId) {
		utils.ErrorNoSuchUser.Write(c)
		return
	}

	r, err := serivces.Feed.Relation(ctx, &rpc.DouyinRelationActionRequest{
		RequestUserId: int64(user),
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

	user, err := jwt.AuthorizedUser(c)
	if err != nil {
		utils.Error(c, err)
		return
	}

	r, err := serivces.Feed.Following(ctx, &rpc.DouyinRelationFollowListRequest{
		UserId:        req.UserId,
		RequestUserId: int64(user),
	})
	if err != nil {
		hlog.Warn(err)
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	list, err := toUserList(r.UserList)
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

	user, err := jwt.AuthorizedUser(c)
	if err != nil {
		utils.Error(c, err)
		return
	}

	r, err := serivces.Feed.Follower(ctx, &rpc.DouyinRelationFollowerListRequest{
		UserId:        req.UserId,
		RequestUserId: int64(user),
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	list, err := toUserList(r.UserList)
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

	user, err := jwt.AuthorizedUser(c)
	if err != nil {
		utils.Error(c, err)
		return
	}

	r, err := serivces.Feed.Friend(context.Background(), &rpc.DouyinRelationFriendListRequest{
		UserId:        req.UserId,
		RequestUserId: int64(user),
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	list, err := toUserList(r.UserList)
	if err != nil {
		utils.Error(c, err)
		return
	}
	resp := &core.DouyinRelationFriendListResponse{
		UserList: list,
	}

	c.JSON(consts.StatusOK, resp)
}
