// This is an I-have-no-idea counter service.
// It tries to achieve something...
//  1. No consistency is provided when database connections fail.
//  2. Final consistency is achieved when an cache item
//     expires and eventually get written into the database.
//  3. It is scalable.
//
// It writes to the database when:
//   - A request tries to update an item that is not in cache.
//   - A value is passed to the `OnExit` callback function
//     (when the item expires or is `Set` a new value).
package db

import (
	"math"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dgraph-io/ristretto"
	"gorm.io/gorm"
)

var db *gorm.DB
var cache *ristretto.Cache

func Init(dialector gorm.Dialector) (err error) {
	db, err = gorm.Open(dialector)
	if err != nil {
		return
	}

	if err = db.AutoMigrate(&CounterDO{}); err != nil {
		return
	}

	cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 10e5,
		MaxCost:     1 << 28,
		BufferItems: 64,
		OnExit: func(val interface{}) {
			if err := val.(*Counters).Exit(); err != nil {
				klog.Error(err)
			}
		},
	})

	return
}

func WithDB(fn func(*gorm.DB)) {
	fn(db)
}

func addOrCreate(id int64, kind int8, delta int32) error {
	if delta == 0 {
		return nil
	}

	model := &CounterDO{}
	result := db.Model(model).
		Where("id = ?", id).
		Where("kind = ?", kind).
		Update("count", gorm.Expr("count + ?", delta))
	if err := result.Error; err != nil {
		return err
	} else if result.RowsAffected == 0 {
		// Slow path, atomic insert or addition.
		model.Id = id
		model.Kind = kind
		model.Count = int32(delta)
		if err := db.Create(model).Error; err != nil {
			// Retry addition.
			if err := db.Model(model).
				Where("id = ?", id).
				Where("kind = ?", kind).
				Update("count", gorm.Expr("count + ?", delta)).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func Increment(id int64, kind int8, delta int32) (err error) {
	value, _ := cache.Get(id)
	if value == nil {
		err = addOrCreate(id, kind, int32(delta))
		return
	}

	counters := value.(*Counters)
	if counters.Lock() {
		counters.Add(kind, delta)
		err = counters.Unlock()
	} else {
		newCounters := &Counters{
			id: id,
		}
		// Approximated.
		for i := 0; i < len(newCounters.counts); i++ {
			newCounters.counts[i] = counters.counts[i] + atomic.LoadInt32(&counters.deltas[i])
		}
		newCounters.Add(kind, delta)
		cache.SetWithTTL(id, newCounters, 0, time.Minute/2)
	}

	return
}

func Get(id int64) (*Counters, error) {
	if c, _ := cache.Get(id); c != nil {
		return c.(*Counters), nil
	}

	counters := &Counters{
		id: id,
	}

	inner := [4]CounterDO{}
	counts := inner[0:0]

	result := db.Where("id = ?", id).Limit(4).Find(&counts)
	if err := result.Error; err != nil {
		return nil, err
	}

	for _, count := range counts {
		counters.counts[count.Kind] = count.Count
	}

	cache.SetWithTTL(id, counters, 0, 5*time.Minute)
	return counters, nil
}

func Invalidate(id int64) {
	cache.Del(id)
}

type Counters struct {
	id     int64
	deltas [4]int32
	counts [4]int32
	// This field records concurrent operations and OnExit calls.
	lock int32
}

// Marks the counters as "OnExit" and ensures that
// the value get synchronized into the database.
func (c *Counters) Exit() error {
	// Inverts the sign bit and increments by 1.
	atomic.AddInt32(&c.lock, math.MinInt32+1)
	return c.Unlock()
}

// Increments the lock holder count.
//
// It returns false if the value should not be used any more
// because it is no longer in the cache.
func (c *Counters) Lock() bool {
	// Increments.
	if atomic.AddInt32(&c.lock, 1) == (math.MinInt32 + 1) {
		return false
	} else {
		return true
	}
}

// Decrements the lock holder count.
//
// If the value is marked as "OnExit",
// the last unlocker will automatically synchronize
// the value into the database.
func (c *Counters) Unlock() (err error) {
	if atomic.AddInt32(&c.lock, -1) == math.MinInt32 {
		for kind := 0; kind < len(c.deltas); kind++ {
			delta := atomic.LoadInt32(&c.deltas[kind])
			e := addOrCreate(c.id, int8(kind), delta)
			if err == nil {
				err = e
			}
		}
	}
	return
}

func (c *Counters) Add(kind int8, delta int32) {
	atomic.AddInt32(&c.deltas[kind], delta)
}

func (c *Counters) Count(kind int8) int32 {
	return c.counts[kind] + atomic.LoadInt32(&c.deltas[kind])
}

type CounterDO struct {
	Id    int64 `gorm:"<-:create;primaryKey;autoIncrement=false"`
	Kind  int8  `gorm:"<-:create;primaryKey;autoIncrement=false"`
	Count int32
}
