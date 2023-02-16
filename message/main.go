package main

import (
	"common"
	rpc "common/kitex_gen/douyin/rpc/messageservice"
	"common/snowy"
	"common/utils"
	"message/internal/cql"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	utils.InitKlog()
	utils.InitEnvVars()

	if err := snowy.Init(); err != nil {
		klog.Fatal(err)
	}

	if err := cql.Init(utils.Env.Cassandra); err != nil {
		klog.Fatal(err)
	}

	svr := rpc.NewServer(
		new(MessageServiceImpl),
		common.WithEtcdOptions(common.MessageServiceName)...,
	)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
