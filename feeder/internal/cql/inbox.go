package cql

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/godruoyi/go-snowflake"
)

const (
	stmt_list_inbox = "SELECT item FROM feed.inbox WHERE user = ? AND item < ?" +
		"ORDER BY item DESC PER PARTITION LIMIT ?"
	stmt_insert_inbox = "INSERT INTO feed.inbox (user, item) VALUES (?, ?)"
)

var snowflakeStart = time.Date(2008, 11, 10, 23, 0, 0, 0, time.UTC)

func ListInbox(user int64, latest time.Time, limit int) ([]int64, error) {
	scanner := session.Query(
		stmt_list_inbox,
		user,
		latest.Sub(snowflakeStart).Milliseconds()<<
			(snowflake.MachineIDLength+snowflake.SequenceLength),
		limit,
	).Iter().Scanner()

	items := make([]int64, 0, limit)
	for scanner.Next() {
		var item int64
		if err := scanner.Scan(&item); err != nil {
			continue
		}
		items = append(items, item)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func PushInboxes(item int64, users []int64) error {
	batch := session.NewBatch(gocql.UnloggedBatch)
	batch.Entries = make([]gocql.BatchEntry, 0, len(users))
	for _, user := range users {
		batch.Entries = append(batch.Entries, gocql.BatchEntry{
			Stmt:       stmt_insert_inbox,
			Args:       []interface{}{user, item},
			Idempotent: true,
		})
	}
	if err := session.ExecuteBatch(batch); err != nil {
		return err
	}
	return nil
}
