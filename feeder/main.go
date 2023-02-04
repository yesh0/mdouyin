package main

import (
	rpc "feeder/kitex_gen/douyin/rpc/feedservice"
	"log"
)

func main() {
	svr := rpc.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
