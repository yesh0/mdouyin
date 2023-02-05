package main

import (
	"log"
	rpc "reaction/kitex_gen/douyin/rpc/reactionservice"
)

func main() {
	svr := rpc.NewServer(new(ReactionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
