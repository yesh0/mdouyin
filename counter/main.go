package main

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"common/utils"
	"counter/db"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	utils.InitKlog()
	utils.InitEnvVars()

	svr := counterservice.NewServer(
		new(CounterServiceImpl),
		common.WithEtcdOptions(common.CounterServiceName)...,
	)

	if err := db.Init(utils.GormDialector()); err != nil {
		klog.Fatal(err)
	}

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
