package main

import (
	"common/kitex_gen/douyin/rpc/reactionservice"
	"common/utils"
	"log"
)

func main() {
	utils.InitKlog()

	svr := reactionservice.NewServer(new(ReactionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
