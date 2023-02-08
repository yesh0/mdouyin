package main

import (
	"common/kitex_gen/douyin/rpc/feedservice"
	"common/snowy"
	"feeder/internal/cql"
	"feeder/internal/db"
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gorm.io/driver/sqlite"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatalln(err)
	}

	if err := cql.Init("127.0.0.1"); err != nil {
		log.Fatalln(err)
	}

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		log.Fatalln(err)
	}

	if err := snowy.Init("127.0.0.1:2379"); err != nil {
		log.Fatalln(err)
	}

	svr := feedservice.NewServer(
		new(FeedServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "feed",
		}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
