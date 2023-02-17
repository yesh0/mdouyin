package db_test

import (
	"common/utils"
	"gateway/internal/db"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	utils.Env.Rdbms = "file::memory:?cache=shared"
	err := db.Init()
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
