package cache

import (
	"context"
	"unsafe"

	"github.com/dgraph-io/ristretto"
	"golang.org/x/time/rate"
)

type Limiter struct {
	rps   int
	burst int
	users *ristretto.Cache
}

func NewRateLimiter(rps int, burst int) (*Limiter, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Limiter{
		rps:   rps,
		burst: burst,
		users: cache,
	}, nil
}

func (l *Limiter) Wait(ctx context.Context, id int64) error {
	i, ok := l.users.Get(id)
	if ok && i != nil {
		return i.(*rate.Limiter).Wait(ctx)
	} else {
		limiter := rate.NewLimiter(rate.Limit(l.rps), 10*l.burst)
		l.users.Set(id, limiter, int64(unsafe.Sizeof(*limiter)))
		return nil
	}
}
