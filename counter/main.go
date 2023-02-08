package main

import (
	"common/kitex_gen/douyin/rpc/counterservice"
	"counter/db"
	"log"

	"gorm.io/driver/sqlite"
)

func main() {
	svr := counterservice.NewServer(new(CounterServiceImpl))

	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		log.Fatalln(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
