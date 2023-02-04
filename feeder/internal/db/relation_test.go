package db_test

import (
	"common/utils"
	"feeder/internal/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutualFollow(t *testing.T) {
	userCount := 200
	for i := 0; i < userCount; i++ {
		for j := i + 1; j < userCount; j++ {
			assert.Equal(t, utils.ErrorOk, db.Follow(int64(i), int64(j)))
			assert.Equal(t, utils.ErrorOk, db.Follow(int64(j), int64(i)))
		}
	}
	max := userCount - 1
	if max > 300 {
		max = 300
	}
	list, err := db.FollowerList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, max)
	list, err = db.FolloweeList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, max)
	list, err = db.FriendList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, max)

	for i := 0; i < userCount; i++ {
		for j := i + 1; j < userCount; j++ {
			assert.Equal(t, utils.ErrorOk, db.Unfollow(int64(i), int64(j)))
			assert.Equal(t, utils.ErrorOk, db.Unfollow(int64(j), int64(i)))
		}
	}
	list, err = db.FollowerList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 0)
	list, err = db.FolloweeList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 0)
	list, err = db.FriendList(5)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 0)
}

func assertRelation(t assert.TestingT, follower, followee int64, follows bool, mutual bool) {
	list, err := db.FollowerList(followee)
	assert.Equal(t, utils.ErrorOk, err)
	if follows {
		assert.Contains(t, list, follower)
	} else {
		assert.NotContains(t, list, follower)
	}

	list, err = db.FolloweeList(follower)
	assert.Equal(t, utils.ErrorOk, err)
	if follows {
		assert.Contains(t, list, followee)
	} else {
		assert.NotContains(t, list, followee)
	}

	list, err = db.FollowerList(follower)
	assert.Equal(t, utils.ErrorOk, err)
	if mutual {
		assert.Contains(t, list, followee)
	} else {
		assert.NotContains(t, list, followee)
	}

	list, err = db.FolloweeList(followee)
	assert.Equal(t, utils.ErrorOk, err)
	if mutual {
		assert.Contains(t, list, follower)
	} else {
		assert.NotContains(t, list, follower)
	}
}

func TestMutualRelation(t *testing.T) {
	assert.Equal(t, utils.ErrorOk, db.Follow(0xf0000, 0xf0001))
	assertRelation(t, 0xf0000, 0xf0001, true, false)
	assert.Equal(t, utils.ErrorOk, db.Follow(0xf0001, 0xf0000))
	assertRelation(t, 0xf0000, 0xf0001, true, true)
	assert.Equal(t, utils.ErrorOk, db.Unfollow(0xf0000, 0xf0001))
	assertRelation(t, 0xf0001, 0xf0000, true, false)
	assert.Equal(t, utils.ErrorOk, db.Unfollow(0xf0001, 0xf0000))
	assertRelation(t, 0xf0001, 0xf0000, false, false)
}