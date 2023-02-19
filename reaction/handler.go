package main

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"common/utils"
	"context"
	"reaction/internal/cache"
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
		if cache.IsFavorite(ctx, req.RequestUserId, req.VideoId) == 1 {
			resp.StatusCode = int32(utils.ErrorRepeatedRequests)
			return
		}
		resp.StatusCode = int32(db.Favorite(req.RequestUserId, req.VideoId))
		if resp.StatusCode == 0 {
			cache.Favorite(ctx, req.RequestUserId, req.VideoId)
			resp.StatusCode = int32(incrementCount(ctx, common.KindVideoFavoriteCount,
				req.VideoId, req.RequestUserId, 1))
		}
	case 2: // Unfavorite
		if cache.IsFavorite(ctx, req.RequestUserId, req.VideoId) == 0 {
			resp.StatusCode = int32(utils.ErrorRepeatedRequests)
			return
		}
		resp.StatusCode = int32(db.Unfavorite(req.RequestUserId, req.VideoId))
		if resp.StatusCode == 0 {
			cache.Unfavorite(ctx, req.RequestUserId, req.VideoId)
			resp.StatusCode = int32(incrementCount(ctx, common.KindVideoFavoriteCount,
				req.VideoId, req.RequestUserId, -1))
		}
	default:
		resp.StatusCode = int32(utils.ErrorWrongParameter)
	}
	return
}

func incrementCount(ctx context.Context, kind int8, video int64, user int64, inc int16) utils.ErrorCode {
	var actions []*rpc.Increment

	if kind == common.KindVideoFavoriteCount {
		actions = []*rpc.Increment{
			{Id: video, Kind: common.KindVideoFavoriteCount, Delta: inc},
			{Id: user, Kind: common.KindUserFavoriteCount, Delta: inc},
		}
	} else {
		actions = []*rpc.Increment{{Id: video, Kind: kind, Delta: inc}}
	}

	_, err := services.Counter.Increment(ctx, &rpc.CounterIncRequest{
		Actions: actions,
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
	for _, video := range favorites[:10] {
		cache.Favorite(ctx, req.UserId, video)
	}
	resp.StatusCode = int32(e)
	if e == utils.ErrorOk {
		resp.VideoList = favorites
	}
	return
}

// TestFavorites implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) TestFavorites(ctx context.Context, req *rpc.FavoriteTestRequest) (resp *rpc.FavoriteTestResponse, err error) {
	resp = rpc.NewFavoriteTestResponse()
	favorites := cache.AreFavorites(ctx, req.RequestUserId, req.Videos)
	remaining := make([]int64, 0)
	for i, fav := range favorites {
		if fav == -1 {
			remaining = append(remaining, req.Videos[i])
		}
	}

	if len(remaining) != 0 {
		other_favs, e := db.IsFavorite(req.RequestUserId, remaining)
		resp.StatusCode = int32(e)
		i := 0
		for j, fav := range favorites {
			if fav == -1 {
				switch other_favs[i] {
				case 0:
					cache.Unfavorite(ctx, req.RequestUserId, req.Videos[j])
				case 1:
					cache.Favorite(ctx, req.RequestUserId, req.Videos[j])
				}
				favorites[j] = other_favs[i]
				i++
			}
			if i >= len(other_favs) {
				break
			}
		}
	}

	if resp.StatusCode == 0 {
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
		resp.StatusCode = int32(incrementCount(ctx, common.KindVideoCommentCount,
			req.VideoId, req.RequestUserId, 1))
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
