package cache

import (
	"gateway/biz/model/douyin/core"
	"unsafe"

	"github.com/dgraph-io/ristretto"
)

var cache *ristretto.Cache

func Init() (err error) {
	cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 10e5,
		MaxCost:     1 << 28,
		BufferItems: 64,
	})
	if err != nil {
		return
	}
	return
}

func GetUser(user int64) *core.User {
	if v, ok := cache.Get(user); ok && v != nil {
		copy := *(v.(*core.User))
		return &copy
	}
	return nil
}

func SetUser(id int64, user *core.User) {
	cache.Set(id, user, int64(unsafe.Sizeof(*user)))
	user.IsFollow = false
	user.Message = nil
	user.MsgType = nil
}

func Flush(user int64) {
	cache.Del(user)
}

func FlushFollowing(user int64) {
	id := uint64(user) | (uint64(1) << 63)
	cache.Del(id)
}

func GetFollowing(user int64) []int64 {
	id := uint64(user) | (uint64(1) << 63)
	if v, ok := cache.Get(id); ok && v != nil {
		return v.([]int64)
	}
	return nil
}

func SetFollowing(user int64, v []int64) {
	id := uint64(user) | (uint64(1) << 63)
	cache.Set(id, v, int64(len(v)*8))
}
