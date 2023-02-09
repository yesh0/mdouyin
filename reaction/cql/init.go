package cql

import "github.com/gocql/gocql"

var script = []string{
	`-- Create a keyspace
	CREATE KEYSPACE IF NOT EXISTS reaction
		WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' }`,
	`-- Create a table
	CREATE TABLE IF NOT EXISTS reaction.comment (
		video BIGINT,
		item BIGINT,
		author BIGINT,
		content TEXT,
		removed TINYINT,
		PRIMARY KEY ((video), item, author),
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
