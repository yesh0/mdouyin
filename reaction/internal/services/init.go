package services

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"common/kitex_gen/douyin/rpc/feedservice"
)

var Counter counterservice.Client
var Feed feedservice.Client

func Init() (err error) {
	Counter, err = counterservice.NewClient(string(common.CounterServiceName), common.WithEtcdResolver())
	if err != nil {
		return
	}

	Feed, err = feedservice.NewClient(string(common.FeederServiceName), common.WithEtcdResolver())
	if err != nil {
		return
	}
	return
}
