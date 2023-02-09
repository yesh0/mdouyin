package services

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"common/kitex_gen/douyin/rpc/feedservice"
	"fmt"
)

var Feed feedservice.Client
var Counter counterservice.Client

func Init() (err error) {
	Feed, err = feedservice.NewClient(string(common.FeederServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create feed client: %s", err)
	}

	Counter, err = counterservice.NewClient(string(common.CounterServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create counter client: %s", err)
	}

	return nil
}
