package services

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"common/kitex_gen/douyin/rpc/feedservice"
	"common/kitex_gen/douyin/rpc/messageservice"
	"common/kitex_gen/douyin/rpc/reactionservice"
	"fmt"
)

var Counter counterservice.Client
var Feed feedservice.Client
var Message messageservice.Client
var Reaction reactionservice.Client

func Init() (err error) {
	Feed, err = feedservice.NewClient(string(common.FeederServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create feed client: %s", err)
	}

	Counter, err = counterservice.NewClient(string(common.CounterServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create counter client: %s", err)
	}

	Message, err = messageservice.NewClient(string(common.MessageServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create message client: %s", err)
	}

	Reaction, err = reactionservice.NewClient(string(common.ReactionServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create reaction client: %s", err)
	}

	return nil
}
