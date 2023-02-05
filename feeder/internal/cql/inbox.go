package cql

import (
	"common/snowy"
	"common/utils"
	"time"

	"github.com/gocql/gocql"
)

const (
	stmt_list_inbox = "SELECT item FROM feed.inbox WHERE user = ? AND item < ?" +
		"ORDER BY item DESC PER PARTITION LIMIT ?"
	stmt_insert_inbox = "INSERT INTO feed.inbox (user, item) VALUES (?, ?)"
)

func ListInbox(user int64, latest time.Time, limit int) ([]int64, utils.ErrorCode) {
	scanner := session.Query(
		stmt_list_inbox,
		user,
		snowy.FromLowerTime(latest),
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
		return nil, utils.ErrorDatabaseError
	}
	return items, utils.ErrorOk
}

func PushInboxes(item int64, users []int64) utils.ErrorCode {
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
		return utils.ErrorDatabaseError
	}
	return utils.ErrorOk
}
