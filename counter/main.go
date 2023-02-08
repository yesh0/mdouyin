package main

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"common/utils"
	"counter/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/sqlite"
)

func main() {
	utils.InitKlog()

	svr := counterservice.NewServer(
		new(CounterServiceImpl),
		common.WithEtcdOptions(common.CounterServiceName)...,
	)

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		klog.Fatal(err)
	}

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
