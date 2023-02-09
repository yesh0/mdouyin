package services

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
)

var Counter counterservice.Client

func Init() (err error) {
	Counter, err = counterservice.NewClient(string(common.CounterServiceName), common.WithEtcdResolver())
	if err != nil {
		return
	}
	return
}
