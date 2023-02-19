package common

import "common/kitex_gen/douyin/rpc"

type RpcServiceName string

const (
	CounterServiceName  RpcServiceName = "counter"
	FeederServiceName   RpcServiceName = "feeder"
	MessageServiceName  RpcServiceName = "message"
	ReactionServiceName RpcServiceName = "reaction"
)

const (
	KindUserFollowerCount  = int8(0)
	KindUserFollowingCount = int8(1)
	KindUserFavoriteCount  = int8(2)
	KindUserWorkCount      = int8(3)
	KindUserTotalFavorited = int8(4)
)

const (
	KindVideoFavoriteCount = int8(0)
	KindVideoCommentCount  = int8(1)
)

func NewIncrement(follower int64, followee int64, unfollow bool) *rpc.CounterIncRequest {
	req := rpc.NewCounterIncRequest()
	var delta int16
	if unfollow {
		delta = -1
	} else {
		delta = 1
	}
	req.Actions = []*rpc.Increment{
		{Id: follower, Kind: KindUserFollowingCount, Delta: delta},
		{Id: followee, Kind: KindUserFollowerCount, Delta: delta},
	}
	return req
}
