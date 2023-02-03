// Code generated by hertz generator.

package core

import (
	"context"

	core "gateway/biz/model/douyin/core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Favorite .
// @router /douyin/favorite/action/ [POST]
func Favorite(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListFavorites .
// @router /douyin/favorite/list/ [GET]
func ListFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Comment .
// @router /douyin/comment/action/ [POST]
func Comment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinCommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListComments .
// @router /douyin/comment/list/ [GET]
func ListComments(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}