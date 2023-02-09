package main

import (
	"common/kitex_gen/douyin/rpc"
	"common/utils"
	"context"
	"reaction/cql"
	"time"
)

// ReactionServiceImpl implements the last service interface defined in the IDL.
type ReactionServiceImpl struct{}

// Favorite implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) Favorite(ctx context.Context, req *rpc.DouyinFavoriteActionRequest) (resp *rpc.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFavorites implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) ListFavorites(ctx context.Context, req *rpc.DouyinFavoriteListRequest) (resp *rpc.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) Comment(ctx context.Context, req *rpc.DouyinCommentActionRequest) (resp *rpc.DouyinCommentActionResponse, err error) {
	resp = rpc.NewDouyinCommentActionResponse()
	switch req.ActionType {
	case 1: // Publishes.
		if req.CommentText == nil || *req.CommentText == "" {
			resp.StatusCode = int32(utils.ErrorWrongInputFormat)
			return
		}
		var code utils.ErrorCode
		resp.Comment, code = cql.AddComment(req.VideoId, req.RequestUserId, *req.CommentText)
		resp.StatusCode = int32(code)
	case 2: // Removes.
		if req.CommentId == nil || *req.CommentId == 0 {
			resp.StatusCode = int32(utils.ErrorWrongInputFormat)
			return
		}
		resp.StatusCode = int32(cql.DeleteComment(req.VideoId, *req.CommentId, req.RequestUserId))
	}
	return
}

// ListComments implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) ListComments(ctx context.Context, req *rpc.DouyinCommentListRequest) (resp *rpc.DouyinCommentListResponse, err error) {
	resp = rpc.NewDouyinCommentListResponse()
	list, e := cql.ListComment(req.VideoId, time.Now(), 300)
	if e != utils.ErrorOk {
		resp.StatusCode = int32(e)
		return
	}
	resp.CommentList = list
	return
}
