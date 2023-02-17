package cache

import (
	"context"
	"encoding/binary"
	"strings"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

const (
	fav_map     = "fav"
	favorited   = "1"
	unfavorited = "0"
)

func Init(host string) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: host,
	})
	_, err = rdb.ClientID(context.Background()).Result()
	return
}

func encode(buf [16]byte, x, y int64) []byte {
	binary.LittleEndian.PutUint64(buf[:8], uint64(x))
	binary.LittleEndian.PutUint64(buf[8:], uint64(y))
	return buf[:]
}

func Favorite(ctx context.Context, user int64, video int64) error {
	buf := [16]byte{}
	return rdb.HSet(ctx, fav_map, encode(buf, user, video), favorited).Err()
}

func Unfavorite(ctx context.Context, user int64, video int64) error {
	buf := [16]byte{}
	return rdb.HSet(ctx, fav_map, encode(buf, user, video), unfavorited).Err()
}

func IsFavorite(ctx context.Context, user int64, video int64) int {
	buf := [16]byte{}
	result, err := rdb.HGet(ctx, fav_map, string(encode(buf, user, video))).Result()
	if err != nil {
		return -1
	}
	return conv(result)
}

func conv(s string) int {
	switch s {
	case favorited:
		return 1
	case unfavorited:
		return 0
	default:
		return -1
	}
}

func AreFavorites(ctx context.Context, user int64, videos []int64) []int8 {
	b := strings.Builder{}
	b.Grow(len(videos) * 16)
	for _, v := range videos {
		buf := [16]byte{}
		b.Write(encode(buf, user, v))
	}
	s := b.String()
	fields := make([]string, 0, len(videos))
	for i := 0; i < len(videos); i++ {
		fields = append(fields, s[i*16:(i+1)*16])
	}
	if result, err := rdb.HMGet(ctx, fav_map, fields...).Result(); err != nil {
		return nil
	} else {
		values := make([]int8, 0, len(videos))
		for _, v := range result {
			var i int
			if v == nil {
				i = -1
			} else if s, ok := v.(string); !ok {
				i = -1
			} else {
				i = conv(s)
			}
			values = append(values, int8(i))
		}
		return values
	}
}
