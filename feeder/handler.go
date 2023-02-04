package main

import (
	"common/utils"
	"context"
	"feeder/internal/db"
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
	switch req.ActionType {
	case 1: // Follow
		if err := db.Follow(req.RequestUserId, req.ToUserId); err != utils.ErrorOk {
			resp.StatusCode = int32(err)
		}
	case 2: // Unfollow
		if err := db.Unfollow(req.RequestUserId, req.ToUserId); err != utils.ErrorOk {
			resp.StatusCode = int32(err)
		}
	default:
		resp.StatusCode = int32(utils.ErrorWrongInputFormat)
	}
	return
}

// Following implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Following(ctx context.Context, req *rpc.DouyinRelationFollowListRequest) (resp *rpc.DouyinRelationFollowListResponse, err error) {
	list, e := db.FolloweeList(req.UserId)
	if err != nil {
		resp.StatusCode = int32(e)
		return
	}
	resp.UserList = list
	return
}

// Follower implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Follower(ctx context.Context, req *rpc.DouyinRelationFollowerListRequest) (resp *rpc.DouyinRelationFollowerListResponse, err error) {
	list, e := db.FollowerList(req.UserId)
	if err != nil {
		resp.StatusCode = int32(e)
		return
	}
	resp.UserList = list
	return
}

// Friend implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Friend(ctx context.Context, req *rpc.DouyinRelationFriendListRequest) (resp *rpc.DouyinRelationFriendListResponse, err error) {
	list, e := db.FriendList(req.UserId)
	if err != nil {
		resp.StatusCode = int32(e)
		return
	}
	resp.UserList = list
	return
}
