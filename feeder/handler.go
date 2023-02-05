package main

import (
	"common/snowy"
	"common/utils"
	"context"
	"feeder/internal/cql"
	"feeder/internal/db"
	"feeder/kitex_gen/douyin/rpc"
	"time"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *rpc.DouyinFeedRequest) (*rpc.DouyinFeedResponse, error) {
	resp := rpc.NewDouyinFeedResponse()
	var videos []db.VideoDO
	var err utils.ErrorCode
	var latest time.Time
	if req.LatestTime == nil {
		latest = time.Now()
	} else {
		latest = time.UnixMilli(*req.LatestTime)
	}

	if req.RequestUserId == 0 {
		videos, err = db.ListLatest(latest, 30)
	} else {
		var ids []int64
		ids, err = cql.ListInbox(req.RequestUserId, latest, 30)
		if err != utils.ErrorOk {
			resp.StatusCode = int32(err)
			return resp, nil
		}
		videos, err = db.FindVideos(ids)
	}
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		return resp, err
	}

	resp.NextTime = new(int64)
	if l := len(videos); l == 0 {
		*resp.NextTime = latest.UnixMilli()
	} else {
		*resp.NextTime = snowy.Time(videos[l-1].Id).UnixMilli()
	}
	resp.VideoList, err = convertList(req.RequestUserId, videos)
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		resp.VideoList = nil
	}
	return resp, nil
}

func extractAuthors(user int64, data []db.VideoDO) (authors map[int64]*rpc.User, err utils.ErrorCode) {
	authors = make(map[int64]*rpc.User)
	ids := make([]int64, 0)
	for _, v := range data {
		if _, ok := authors[v.Author]; ok {
			continue
		}
		ids = append(ids, v.Author)
		authors[v.Author] = &rpc.User{Id: v.Author}
	}
	if user != 0 {
		var followees []int64
		followees, err = db.FilterFollowees(user, ids)
		for _, followee := range followees {
			authors[followee].IsFollow = true
		}
	}
	return
}

func convertList(user int64, data []db.VideoDO) (videos []*rpc.Video, err utils.ErrorCode) {
	var authors map[int64]*rpc.User
	authors, err = extractAuthors(user, data)
	videos = make([]*rpc.Video, 0, len(videos))
	for _, video := range data {
		videos = append(videos, &rpc.Video{
			Id:       video.Id,
			Author:   authors[video.Author],
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			Title:    video.Title,
		})
	}
	return
}

// Publish implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Publish(ctx context.Context, req *rpc.DouyinPublishActionRequest) (*rpc.DouyinPublishActionResponse, error) {
	resp := rpc.NewDouyinPublishActionResponse()
	if req.RequestUserId != req.Video.Author.Id {
		resp.StatusCode = int32(utils.ErrorUnauthorized)
		return resp, nil
	}

	var id int64
	id, err := db.InsertVideo(db.VideoDO{
		Author:   req.Video.Author.Id,
		PlayUrl:  req.Video.PlayUrl,
		CoverUrl: req.Video.CoverUrl,
		Title:    req.Video.Title,
	})
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		return resp, nil
	}

	// TODO: Paged fetch
	var followers []int64
	followers, err = db.FollowerList(req.Video.Author.Id)
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		return resp, nil
	}

	resp.StatusCode = int32(cql.PushInboxes(id, followers))
	return resp, nil
}

// List implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) List(ctx context.Context, req *rpc.DouyinPublishListRequest) (*rpc.DouyinPublishListResponse, error) {
	resp := rpc.NewDouyinPublishListResponse()
	videos, err := db.ListVideos(req.UserId, 300)
	if err != utils.ErrorOk {
		resp.StatusCode = int32(utils.ErrorUnanticipated)
		return resp, nil
	}
	resp.VideoList, err = convertList(req.RequestUserId, videos)
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		resp.VideoList = nil
	}
	return resp, nil
}

// Relation implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Relation(ctx context.Context, req *rpc.DouyinRelationActionRequest) (resp *rpc.DouyinRelationActionResponse, err error) {
	resp = rpc.NewDouyinRelationActionResponse()
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
	resp = rpc.NewDouyinRelationFollowListResponse()
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
	resp = rpc.NewDouyinRelationFollowerListResponse()
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
	resp = rpc.NewDouyinRelationFriendListResponse()
	list, e := db.FriendList(req.UserId)
	if err != nil {
		resp.StatusCode = int32(e)
		return
	}
	resp.UserList = list
	return
}
