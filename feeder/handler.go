package main

import (
	"context"
	rpc "feeder/kitex_gen/douyin/rpc"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *rpc.DouyinFeedRequest) (resp *rpc.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// Publish implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Publish(ctx context.Context, req *rpc.DouyinPublishActionRequest) (resp *rpc.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) List(ctx context.Context, req *rpc.DouyinPublishListRequest) (resp *rpc.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// Relation implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Relation(ctx context.Context, req *rpc.DouyinRelationActionRequest) (resp *rpc.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// Following implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Following(ctx context.Context, req *rpc.DouyinRelationFollowListRequest) (resp *rpc.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// Follower implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Follower(ctx context.Context, req *rpc.DouyinRelationFollowerListRequest) (resp *rpc.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// Friend implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Friend(ctx context.Context, req *rpc.DouyinRelationFriendListRequest) (resp *rpc.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
