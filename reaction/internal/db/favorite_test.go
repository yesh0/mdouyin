package db_test

import (
	"common/utils"
	"log"
	"math/rand"
	"reaction/internal/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestMain(m *testing.M) {
	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func TestFavorites(t *testing.T) {
	favorites, err := db.ListFavorites(1, 300)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Empty(t, favorites)

	rand.Seed(time.Now().Unix())
	ids := make(map[int64]struct{})
	for i := 0; i < 240; i++ {
		id := rand.Int63()
		assert.Equal(t, utils.ErrorOk, db.Favorite(1, id))
		ids[id] = struct{}{}
	}

	favorites, err = db.ListFavorites(1, 300)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, favorites, len(ids))
	for _, fav := range favorites {
		_, ok := ids[fav]
		assert.True(t, ok)
	}

	list := make([]int64, 0)
	for id := range ids {
		list = append(list, id)
	}
	for i := 0; i < 240; i++ {
		list = append(list, rand.Int63())
	}
	isFavorites, err := db.IsFavorite(1, list)
	assert.Equal(t, utils.ErrorOk, err)
	for i, v := range isFavorites {
		_, ok := ids[list[i]]
		if v == 0 {
			assert.False(t, ok)
		} else {
			assert.True(t, ok)
		}
	}

	for id := range ids {
		assert.Equal(t, utils.ErrorOk, db.Unfavorite(1, id))
	}
	favorites, err = db.ListFavorites(1, 300)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Empty(t, favorites)
}
