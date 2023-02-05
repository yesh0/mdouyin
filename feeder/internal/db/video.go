package db

import (
	"common/snowy"
	"common/utils"
	"time"

	"github.com/godruoyi/go-snowflake"
)

type VideoDO struct {
	Id       int64 `gorm:"<-:create;primaryKey;autoIncrement=false"`
	Author   int64 `gorm:"index"`
	PlayUrl  string
	CoverUrl string
	Title    string
}

func InsertVideo(video VideoDO) (int64, utils.ErrorCode) {
	video.Id = int64(snowflake.ID())
	if err := db.Create(&video).Error; err != nil {
		return 0, utils.ErrorDatabaseError
	}
	return video.Id, utils.ErrorOk
}

func ListVideos(author int64, limit int) ([]VideoDO, utils.ErrorCode) {
	var videos []VideoDO
	if err := db.Where("author", author).
		Order("id DESC").Limit(limit).
		Find(&videos).Error; err != nil {
		return nil, utils.ErrorDatabaseError
	}
	return videos, utils.ErrorOk
}

func ListLatest(latest time.Time, limit int) ([]VideoDO, utils.ErrorCode) {
	var videos []VideoDO
	if err := db.Where("id < ?", snowy.FromLowerTime(latest)).Order("id DESC").Limit(limit).
		Find(&videos).Error; err != nil {
		return nil, utils.ErrorDatabaseError
	}
	return videos, utils.ErrorOk
}

func FindVideos(ids []int64) ([]VideoDO, utils.ErrorCode) {
	videos := make([]VideoDO, 0, len(ids))
	if err := db.Find(&videos, ids).Error; err != nil {
		return nil, utils.ErrorDatabaseError
	}
	return videos, utils.ErrorOk
}
