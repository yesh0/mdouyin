package main

import (
	"common"
	"common/kitex_gen/douyin/rpc/reactionservice"
	"common/snowy"
	"common/utils"
	"log"
	"reaction/internal/cql"
	"reaction/internal/db"
	"reaction/internal/services"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	utils.InitKlog()
	utils.InitEnvVars()

	svr := reactionservice.NewServer(
		new(ReactionServiceImpl),
		common.WithEtcdOptions(common.ReactionServiceName)...,
	)

	if err := cql.Init(utils.Env.Cassandra); err != nil {
		klog.Fatal(err)
	}

	if err := snowy.Init(); err != nil {
		klog.Fatal(err)
	}

	if err := services.Init(); err != nil {
		klog.Fatal(err)
	}

	if err := db.Init(utils.GormDialector()); err != nil {
		klog.Fatal(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
