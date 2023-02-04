package db_test

import (
	"feeder/internal/db"
	"log"
	"testing"

	"gorm.io/driver/sqlite"
)

func TestMain(m *testing.M) {
	if err := db.Init(sqlite.Open("file::memory:?cache=shared")); err != nil {
		log.Fatalln(err)
	}
	m.Run()
}
