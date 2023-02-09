package main

import (
	"common"
	"common/kitex_gen/douyin/rpc/reactionservice"
	"common/utils"
	"log"
	"reaction/internal/cql"
	"reaction/internal/services"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	utils.InitKlog()

	svr := reactionservice.NewServer(
		new(ReactionServiceImpl),
		common.WithEtcdOptions(common.ReactionServiceName)...,
	)

	if err := cql.Init("127.0.0.1"); err != nil {
		klog.Fatal(err)
	}

	if err := services.Init(); err != nil {
		klog.Fatal(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
