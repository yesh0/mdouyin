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

	if err := snowy.Init("127.0.0.1:2379"); err != nil {
		klog.Fatal(err)
	}

	if err := cql.Init("127.0.0.1"); err != nil {
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
