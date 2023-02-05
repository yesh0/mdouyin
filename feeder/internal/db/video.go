package db

import (
	"common/snowy"
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

func InsertVideo(video VideoDO) (int64, error) {
	video.Id = int64(snowflake.ID())
	if err := db.Create(&video).Error; err != nil {
		return 0, err
	}
	return video.Id, nil
}

func ListVideos(author int64, limit int) ([]VideoDO, error) {
	var videos []VideoDO
	if err := db.Where("author", author).
		Order("id DESC").Limit(limit).
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func ListLatest(latest time.Time, limit int) ([]VideoDO, error) {
	var videos []VideoDO
	if err := db.Where("id < ?", snowy.FromLowerTime(latest)).Order("id DESC").Limit(limit).
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func FindVideos(ids []int64) ([]VideoDO, error) {
	videos := make([]VideoDO, 0, len(ids))
	if err := db.Find(&videos, ids).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
