package main

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"common/utils"
	"context"
	"reaction/internal/cql"
	"reaction/internal/db"
	"reaction/internal/services"
	"time"
)

// ReactionServiceImpl implements the last service interface defined in the IDL.
type ReactionServiceImpl struct{}

// Favorite implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) Favorite(ctx context.Context, req *rpc.DouyinFavoriteActionRequest) (resp *rpc.DouyinFavoriteActionResponse, err error) {
	resp = rpc.NewDouyinFavoriteActionResponse()
	switch req.ActionType {
	case 1: // Favorite
		resp.StatusCode = int32(db.Favorite(req.RequestUserId, req.VideoId))
		if resp.StatusCode == 0 {
			resp.StatusCode = int32(incrementCount(ctx, req.VideoId, 1))
		}
	case 2: // Unfavorite
		resp.StatusCode = int32(db.Unfavorite(req.RequestUserId, req.VideoId))
		if resp.StatusCode == 0 {
			resp.StatusCode = int32(incrementCount(ctx, req.VideoId, -1))
		}
	default:
		resp.StatusCode = int32(utils.ErrorWrongParameter)
	}
	return
}

func incrementCount(ctx context.Context, video int64, inc int16) utils.ErrorCode {
	_, err := services.Counter.Increment(ctx, &rpc.CounterIncRequest{
		Actions: []*rpc.Increment{
			{
				Id:    video,
				Kind:  common.KindVideoFavoriteCount,
				Delta: inc,
			},
		},
	})
	if err != nil {
		return utils.ErrorRpcTimeout
	} else {
		return utils.ErrorOk
	}
}

// ListFavorites implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) ListFavorites(ctx context.Context, req *rpc.DouyinFavoriteListRequest) (resp *rpc.DouyinFavoriteListResponse, err error) {
	resp = rpc.NewDouyinFavoriteListResponse()
	favorites, e := db.ListFavorites(req.UserId, 300)
	resp.StatusCode = int32(e)
	if e == utils.ErrorOk {
		resp.VideoList = favorites
	}
	return
}

// TestFavorites implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) TestFavorites(ctx context.Context, req *rpc.FavoriteTestRequest) (resp *rpc.FavoriteTestResponse, err error) {
	resp = rpc.NewFavoriteTestResponse()
	favorites, e := db.IsFavorite(req.RequestUserId, req.Videos)
	resp.StatusCode = int32(e)
	if e == utils.ErrorOk {
		resp.IsFavorites = favorites
	}
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
		if code != utils.ErrorOk {
			resp.StatusCode = int32(code)
			return
		}
		services.Counter.Increment(ctx, &rpc.CounterIncRequest{
			Actions: []*rpc.Increment{
				{
					Id:    resp.Comment.Id,
					Kind:  common.KindVideoCommentCount,
					Delta: 1,
				},
			},
		})
	case 2: // Removes.
		if req.CommentId == nil || *req.CommentId == 0 {
			resp.StatusCode = int32(utils.ErrorWrongInputFormat)
			return
		}
		resp.StatusCode = int32(cql.DeleteComment(req.VideoId, *req.CommentId, req.RequestUserId))
		// We are not decrementing the counts. No.
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
