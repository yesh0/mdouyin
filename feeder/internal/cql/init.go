package cql

import "github.com/gocql/gocql"

var script = []string{
	`-- Create a keyspace
	CREATE KEYSPACE IF NOT EXISTS feed
		WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' }`,
	`-- Create a table
	CREATE TABLE IF NOT EXISTS feed.inbox (
		user BIGINT,
		item BIGINT,
		PRIMARY KEY ((user), item),
	)`,
}

var session *gocql.Session

func Init(instances ...string) error {
	cluster := gocql.NewCluster(instances...)
	cluster.Consistency = gocql.LocalQuorum
	cluster.NumConns = 3

	var err error
	session, err = cluster.CreateSession()
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
