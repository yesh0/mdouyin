package cql_test

import (
	"feeder/internal/cql"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	if err := cql.Init("127.0.0.1"); err != nil {
		log.Fatalln("the tests require a locally running cassandra instance", err)
	}
	m.Run()
}
