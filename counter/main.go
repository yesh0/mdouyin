package main

import (
	"counter/db"
	rpc "counter/kitex_gen/douyin/rpc/counterservice"
	"log"

	"gorm.io/driver/sqlite"
)

func main() {
	svr := rpc.NewServer(new(CounterServiceImpl))

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		log.Fatalln(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
