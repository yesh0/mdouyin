package db_test

import (
	"common/utils"
	"counter/db"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	utils.Env.Rdbms = "file::memory:?cache=shared"
	if err := db.Init(utils.GormDialector()); err != nil {
		log.Fatalln(err)
	}
	db.WithDB(func(d *gorm.DB) {
		sqlDB, err := d.DB()
		if err != nil {
			log.Fatalln(err)
		}
		sqlDB.SetMaxOpenConns(1)
	})

	m.Run()
}

func TestSingleThreaded(t *testing.T) {
	id := int64(1)
	db.Increment(id, 0, 5)
	c, err := db.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, int32(5), c.Count(0))
}

func TestSimple(t *testing.T) {
	counters, err := db.Get(2)
	// Ristretto batches, so we need to wait until things apply.
	// Or else, the first addition might go straight to the DB,
	// leaving a little inconsistency out there.
	assert.Nil(t, err)
	time.Sleep(time.Second)

	assert.NotNil(t, counters)
	assert.Equal(t, int32(0), counters.Count(0))
	assert.Nil(t, db.Increment(2, 0, 1))

	counters, err = db.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, counters)
	assert.Equal(t, int32(1), counters.Count(0))
}

func TestCounter(t *testing.T) {
	rand.Seed(time.Now().Unix())
	id := rand.Int63()

	concurrency := 10000
	kinds := 4
	result, err := db.Get(id)
	assert.Nil(t, err)
	for i := 0; i < kinds; i++ {
		assert.Zero(t, result.Count(int8(i)))
	}
	// Ristretto batches, so we need to wait until things apply.
	// Or else, the first addition might go straight to the DB,
	// leaving a little inconsistency out there.
	time.Sleep(time.Second)

	group := sync.WaitGroup{}
	group.Add(kinds * concurrency)
	for i := 0; i < concurrency; i++ {
		for j := 0; j < kinds; j++ {
			final_j := j
			go func() {
				for i := 1; i < 1001; i++ {
					db.Increment(id, int8(final_j), int32(i))
				}
				group.Done()
			}()
		}
	}

	group.Wait()
	result, err = db.Get(id)
	assert.Nil(t, err)
	for i := 0; i < kinds; i++ {
		println(result.Count(0))
		assert.Equal(t, int32(500500*concurrency), result.Count(int8(i)))
	}
	db.Invalidate(id)
	result, err = db.Get(id)
	assert.Nil(t, err)
	for i := 0; i < kinds; i++ {
		assert.Equal(t, int32(500500*concurrency), result.Count(int8(i)))
	}
}
