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
	"gorm.io/driver/sqlite"
)

func main() {
	utils.InitKlog()

	if err := cql.Init("127.0.0.1"); err != nil {
		klog.Fatal(err)
	}

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		klog.Fatal(err)
	}

	if err := snowy.Init("127.0.0.1:2379"); err != nil {
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
