package main

import (
	"common"
	"common/kitex_gen/douyin/rpc/feedservice"
	"common/snowy"
	"common/utils"
	"feeder/internal/cql"
	"feeder/internal/db"
	"feeder/internal/services"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	utils.InitKlog()
	utils.InitEnvVars()

	if err := cql.Init(utils.Env.Cassandra); err != nil {
		klog.Fatal(err)
	}

	if err := db.Init(utils.GormDialector()); err != nil {
		klog.Fatal(err)
	}

	if err := snowy.Init(); err != nil {
		klog.Fatal(err)
	}

	if err := services.Init(); err != nil {
		klog.Fatal(err)
	}

	svr := feedservice.NewServer(
		new(FeedServiceImpl),
		common.WithEtcdOptions(common.FeederServiceName)...,
	)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
