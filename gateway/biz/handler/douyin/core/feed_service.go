// Code generated by hertz generator.

package core

import (
	"common/kitex_gen/douyin/rpc"
	"common/utils"
	"context"
	"os"
	"path"

	core "gateway/biz/model/douyin/core"
	"gateway/internal/db"
	"gateway/internal/jwt"
	"gateway/internal/services"
	"gateway/internal/videos"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	id, err := jwt.AuthorizedUser(c)
	if err != nil {
		id = 0
	}

	r, err := services.Feed.Feed(ctx, &rpc.DouyinFeedRequest{
		LatestTime:    req.LatestTime,
		RequestUserId: id,
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if utils.RpcError(c, r.StatusCode) {
		return
	}

	resp := new(core.DouyinFeedResponse)
	resp.NextTime = r.NextTime
	resp.VideoList, err = generateVideoList(r.VideoList)
	if err != nil {
		utils.Error(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

func generateVideoList(info []*rpc.Video) (vs []*core.Video, err error) {
	vs = make([]*core.Video, 0, len(info))

	authorIds := make([]int64, 0)
	authors := make(map[int64]*core.User)
	unknown := fromUser(&db.UserDO{
		Id:   0,
		Name: "Suspended User",
	})
	unknownFollowed := fromUser(&db.UserDO{
		Id:   0,
		Name: "Suspended User",
	})
	unknownFollowed.IsFollow = true
	for _, v := range info {
		if _, ok := authors[v.Author.Id]; ok {
			continue
		}
		if v.Author.IsFollow {
			authors[v.Author.Id] = unknownFollowed
		} else {
			authors[v.Author.Id] = unknown
		}
		authorIds = append(authorIds, v.Author.Id)
	}
	authorLookups, err := db.FindUsersByIds(authorIds)
	for _, v := range authorLookups {
		followed := authors[v.Id].IsFollow
		authors[v.Id] = fromUser(&v)
		if followed {
			authors[v.Id].IsFollow = true
		}
	}

	for _, video := range info {
		// TODO: Fetch counts
		vs = append(vs, &core.Video{
			Id:       video.Id,
			Author:   authors[video.Author.Id],
			PlayUrl:  videos.BaseUrl() + path.Join("/media", video.PlayUrl),
			CoverUrl: videos.BaseUrl() + path.Join("/media", video.CoverUrl),
			Title:    video.Title,
		})
	}
	return
}

// Publish .
// @router /douyin/publish/action/ [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	var err error

	// Videos saved as temporary files automatically by Hertz.
	form, err := c.MultipartForm()
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	// We have to parse the form ourselves.
	tokenFields := form.Value["token"]
	if len(tokenFields) == 0 {
		utils.ErrorUnauthorized.Write(c)
		return
	}
	titleFields := form.Value["title"]
	title := titleFields[0]
	if len(titleFields) == 0 {
		utils.ErrorWrongInputFormat.With("expecting a title").Write(c)
		return
	}
	fileFields := form.File["data"]
	if len(fileFields) == 0 {
		utils.ErrorWrongInputFormat.With("expecting a video").Write(c)
		return
	}

	id, _, err := jwt.Validate(tokenFields[0])
	if err != nil {
		utils.Error(c, err)
		return
	}

	// Check the file
	file := fileFields[0]
	if err := videos.CheckMagic(file); err != nil {
		utils.Error(c, err)
		return
	}

	relPath, err := videos.NewLocalVideo()
	if err != nil {
		utils.Error(c, err)
		return
	}
	fullPath := videos.Storage(relPath)
	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		hlog.Error("error saving uploaded file", err)
		utils.ErrorFilesystem.Write(c)
		return
	}

	// TODO: Queue and rate limit
	go func() {
		if err := videos.ValidateVideo(fullPath); err != nil {
			if err := os.Remove(fullPath); err != nil {
				hlog.Error("unable to remove file", err)
			}
			return
		}
		cover := relPath + ".png"
		if err := videos.GenerateCover(fullPath, videos.Storage(cover)); err != nil {
			hlog.Error("unable to generate cover", err)
			if err := os.Remove(fullPath); err != nil {
				hlog.Error("unable to remove file", err)
			}
			return
		}

		r, err := services.Feed.Publish(ctx, &rpc.DouyinPublishActionRequest{
			RequestUserId: id,
			Video: &rpc.Video{
				Author:   &rpc.User{Id: id},
				PlayUrl:  relPath,
				CoverUrl: cover,
				Title:    title,
			},
		})
		if err != nil {
			hlog.Warn("rpc failed", err)
		}
		if !utils.ErrorOk.IsCode(r.StatusCode) {
			hlog.Warn("rpc failed", utils.ErrorCode(r.StatusCode).Error())
		}
	}()

	resp := new(core.DouyinPublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// List .
// @router /douyin/publish/action/ [GET]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.InvalidInput(c, err)
		return
	}

	r, err := services.Feed.List(ctx, &rpc.DouyinPublishListRequest{
		UserId:        req.UserId,
		RequestUserId: 0,
	})
	if err != nil {
		utils.ErrorRpcTimeout.Write(c)
		return
	}
	if !utils.ErrorOk.IsCode(r.StatusCode) {
		utils.ErrorCode(r.StatusCode).Write(c)
		return
	}

	videos, err := generateVideoList(r.VideoList)
	if err != nil {
		utils.Error(c, err)
	}

	resp := &core.DouyinPublishListResponse{
		VideoList: videos,
	}

	c.JSON(consts.StatusOK, resp)
}
