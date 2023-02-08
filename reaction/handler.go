package main

import (
	"common/kitex_gen/douyin/rpc"
	"context"
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
	// TODO: Your code here...
	return
}

// ListComments implements the ReactionServiceImpl interface.
func (s *ReactionServiceImpl) ListComments(ctx context.Context, req *rpc.DouyinCommentListRequest) (resp *rpc.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
