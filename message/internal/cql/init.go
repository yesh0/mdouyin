package cql

import (
	"common/utils"

	"github.com/gocql/gocql"
)

var script = []string{
	`-- Create a keyspace
	CREATE KEYSPACE IF NOT EXISTS message
		WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' }`,
	`-- Create a table
	CREATE TABLE IF NOT EXISTS message.conversation (
		first BIGINT,
		second BIGINT,
		id BIGINT,
		status TINYINT,
		message TEXT,
		PRIMARY KEY ((first, second), id),
	)`,
}

var session *gocql.Session

func Init(instances ...string) error {
	cluster := gocql.NewCluster(instances...)
	cluster.Consistency = gocql.LocalQuorum
	cluster.NumConns = 3

	var err error
	session, err = utils.GetSession(cluster)
	if err != nil {
		return err
	}

	for _, line := range script {
		if err := session.Query(line).Exec(); err != nil {
			return err
		}
	}

	return nil
}
