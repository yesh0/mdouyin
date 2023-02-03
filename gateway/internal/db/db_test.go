package db_test

import (
	"gateway/internal/db"
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
)

func TestMain(m *testing.M) {
	err := db.Init(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
