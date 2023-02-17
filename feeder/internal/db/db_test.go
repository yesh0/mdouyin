package db_test

import (
	"common/utils"
	"feeder/internal/db"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	utils.Env.Rdbms = "file::memory:?cache=shared"
	if err := db.Init(); err != nil {
		log.Fatalln(err)
	}
	m.Run()
}
