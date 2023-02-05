package db_test

import (
	"common/utils"
	"feeder/internal/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVideoListing(t *testing.T) {
	id1, err := db.InsertVideo(db.VideoDO{
		Author:   1,
		PlayUrl:  "/url1",
		CoverUrl: "/url1.png",
		Title:    "Title1",
	})
	assert.Equal(t, utils.ErrorOk, err)

	list, err := db.ListLatest(time.Now(), 1)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 1)
	assert.Equal(t, int64(id1), list[0].Id)
	list, err = db.ListVideos(1, 1)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 1)
	assert.Equal(t, int64(id1), list[0].Id)

	id2, err := db.InsertVideo(db.VideoDO{
		Author:   1,
		PlayUrl:  "/url1",
		CoverUrl: "/url1.png",
		Title:    "Title1",
	})
	assert.Equal(t, utils.ErrorOk, err)
	assert.Greater(t, id2, id1)

	list, err = db.ListLatest(time.Now(), 1)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 1)
	assert.Equal(t, int64(id2), list[0].Id)
	list, err = db.ListVideos(1, 1)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 1)
	assert.Equal(t, int64(id2), list[0].Id)

	videos, err := db.FindVideos([]int64{id1, id2})
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, videos, 2)
	assert.Equal(t, id1, videos[0].Id)
	assert.Equal(t, id2, videos[1].Id)
}
