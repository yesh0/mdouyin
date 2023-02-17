package cache_test

import (
	"context"
	"log"
	"math/rand"
	"reaction/internal/cache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := cache.Init("localhost:6379"); err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func TestRedis(t *testing.T) {
	rand.Seed(time.Now().Unix())
	u := rand.Int63()
	v := rand.Int63()
	assert.Equal(t, -1, cache.IsFavorite(context.Background(), u, v))
	assert.Nil(t, cache.Favorite(context.Background(), u, v))
	assert.Equal(t, 1, cache.IsFavorite(context.Background(), u, v))
	assert.Nil(t, cache.Unfavorite(context.Background(), u, v))
	assert.Equal(t, 0, cache.IsFavorite(context.Background(), u, v))

	favs := cache.AreFavorites(context.Background(), u, []int64{u, v})
	assert.Len(t, favs, 2)
	assert.Equal(t, int(-1), favs[0])
	assert.Equal(t, int(0), favs[1])
}
