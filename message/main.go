package main

import (
	rpc "common/kitex_gen/douyin/rpc/messageservice"
	"log"
)

func main() {
	svr := rpc.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
