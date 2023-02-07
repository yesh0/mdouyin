package main

import (
	"context"
	"counter/db"
	rpc "counter/kitex_gen/douyin/rpc"
	"log"
)

// CounterServiceImpl implements the last service interface defined in the IDL.
type CounterServiceImpl struct{}

// Increment implements the CounterServiceImpl interface.
func (s *CounterServiceImpl) Increment(ctx context.Context, req *rpc.CounterIncRequest) (resp *rpc.CounterNopResponse, err error) {
	for _, action := range req.Actions {
		if err := db.Increment(action.Id, action.Kind, int32(action.Delta)); err != nil {
			log.Println(err)
		}
	}
	resp = rpc.NewCounterNopResponse()
	return
}

// Fetch implements the CounterServiceImpl interface.
func (s *CounterServiceImpl) Fetch(ctx context.Context, req *rpc.CounterGetRequest) (resp *rpc.CounterGetResponse, err error) {
	resp = rpc.NewCounterGetResponse()
	resp.Counters = make([]*rpc.Counts, 0, len(req.Id))
	for _, id := range req.Id {
		counts := make([]*rpc.KindCount, 0, len(req.Kinds))
		if count, err := db.Get(id); err != nil {
			log.Println(err)
		} else {
			for _, kind := range req.Kinds {
				counts = append(counts, &rpc.KindCount{
					Kind:  kind,
					Count: count.Count(kind),
				})
			}
		}
		resp.Counters = append(resp.Counters, &rpc.Counts{
			Id:         id,
			KindCounts: counts,
		})
	}
	return
}
