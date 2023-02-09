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
	"gorm.io/driver/sqlite"
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

	if err := snowy.Init("127.0.0.1:2379"); err != nil {
		klog.Fatal(err)
	}

	if err := services.Init(); err != nil {
		klog.Fatal(err)
	}

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		klog.Fatal(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
