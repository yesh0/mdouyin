package main

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"common/snowy"
	"common/utils"
	"context"
	"feeder/internal/cql"
	"feeder/internal/db"
	"feeder/internal/services"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *rpc.DouyinFeedRequest) (*rpc.DouyinFeedResponse, error) {
	resp := rpc.NewDouyinFeedResponse()
	var videos []db.VideoDO
	var err utils.ErrorCode
	var latest time.Time
	if req.LatestTime == nil || *req.LatestTime == 0 {
		latest = time.Now()
	} else {
		latest = time.UnixMilli(*req.LatestTime)
	}

	var ids []int64
	if req.RequestUserId == 0 {
		videos, err = db.ListLatest(latest, 30)
	} else {
		ids, err = cql.ListInbox(req.RequestUserId, latest, 30)
		if err != utils.ErrorOk {
			resp.StatusCode = int32(err)
			return resp, nil
		}
		videos, err = db.FindVideos(ids)
	}
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		return resp, nil
	}

	resp.NextTime = new(int64)
	if l := len(videos); l == 0 {
		*resp.NextTime = latest.UnixMilli()
	} else {
		*resp.NextTime = snowy.Time(videos[l-1].Id).UnixMilli()
	}
	resp.VideoList, err = convertList(ctx, req.RequestUserId, ids, videos)
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

func fillVideoCounts(ctx context.Context, ids []int64, videos []*rpc.Video) utils.ErrorCode {
	r, err := services.Counter.Fetch(ctx, &rpc.CounterGetRequest{
		Id:    ids,
		Kinds: []int8{common.KindVideoFavoriteCount, common.KindVideoCommentCount},
	})
	if err != nil {
		return utils.ErrorRpcTimeout
	}

	mapping := make(map[int64]int)
	for i, v := range videos {
		mapping[v.Id] = i
	}

	for _, counter := range r.Counters {
		i, ok := mapping[counter.Id]
		if !ok {
			klog.Warnf("extra counters: %v %v %v", ids, videos, r.Counters)
			continue
		}
		for _, c := range counter.KindCounts {
			switch c.Kind {
			case common.KindVideoFavoriteCount:
				videos[i].FavoriteCount = int64(c.Count)
			case common.KindVideoCommentCount:
				videos[i].CommentCount = int64(c.Count)
			}
		}
	}
	return utils.ErrorOk
}

func convertList(ctx context.Context, user int64,
	ids []int64, data []db.VideoDO) (videos []*rpc.Video, err utils.ErrorCode) {
	if ids == nil {
		if data == nil {
			err = utils.ErrorInternalError
			return
		}
		ids = make([]int64, 0, len(data))
		for _, v := range data {
			ids = append(ids, v.Id)
		}
	}
	if data == nil {
		data, err = db.FindVideos(ids)
		if err != utils.ErrorOk {
			return
		}
	}

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
	if err == utils.ErrorOk {
		err = fillVideoCounts(ctx, ids, videos)
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
	resp.VideoList, err = convertList(ctx, req.RequestUserId, nil, videos)
	if err != utils.ErrorOk {
		resp.StatusCode = int32(err)
		resp.VideoList = nil
	}
	return resp, nil
}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) VideoInfo(ctx context.Context, req *rpc.VideoBatchInfoRequest) (*rpc.VideoBatchInfoResponse, error) {
	resp := rpc.NewVideoBatchInfoResponse()
	converted, err := convertList(ctx, req.RequestUserId, req.VideoIds, nil)
	if err != utils.ErrorOk {
		return resp, nil
	}
	resp.Videos = converted
	return resp, nil
}

// Relation implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Relation(ctx context.Context, req *rpc.DouyinRelationActionRequest) (resp *rpc.DouyinRelationActionResponse, err error) {
	resp = rpc.NewDouyinRelationActionResponse()
	switch req.ActionType {
	case 1: // Follow
		if err := db.Follow(req.RequestUserId, req.ToUserId); err != utils.ErrorOk {
			resp.StatusCode = int32(err)
		} else {
			_, err := services.Counter.Increment(ctx,
				common.NewIncrement(req.RequestUserId, req.ToUserId, false))
			if err != nil {
				resp.StatusCode = int32(utils.ErrorRpcTimeout)
			}
		}
	case 2: // Unfollow
		if err := db.Unfollow(req.RequestUserId, req.ToUserId); err != utils.ErrorOk {
			resp.StatusCode = int32(err)
		} else {
			_, err := services.Counter.Increment(ctx,
				common.NewIncrement(req.RequestUserId, req.ToUserId, true))
			if err != nil {
				resp.StatusCode = int32(utils.ErrorRpcTimeout)
			}
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

// IsFriend implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) IsFriend(ctx context.Context, req *rpc.FriendCheckRequest) (resp *rpc.FriendCheckResponse, err error) {
	resp = rpc.NewFriendCheckResponse()
	if db.IsMutual(req.RequestUserId, req.UserId) == utils.ErrorOk {
		resp.IsFriend = 1
	}
	return
}
