// Code generated by hertz generator.

package core

import (
	"context"

	core "gateway/biz/model/douyin/core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Relation .
// @router /douyin/relation/action/ [POST]
func Relation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
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
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Follower .
// @router /douyin/relation/follower/list/ [GET]
func Follower(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinRelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Friend .
// @router /douyin/relation/friend/list/ [GET]
func Friend(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}