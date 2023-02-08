package main

import (
	"common/kitex_gen/douyin/rpc/reactionservice"
	"log"
)

func main() {
	svr := reactionservice.NewServer(new(ReactionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
