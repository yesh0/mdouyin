package cql

import (
	"common/kitex_gen/douyin/rpc"
	"common/snowy"
	"common/utils"
	"context"
	"sync"
	"time"

	"github.com/gocql/gocql"
	"golang.org/x/sync/semaphore"
)

const (
	stmt_list_latest_conversation = "SELECT id, status, message FROM message.conversation " +
		"WHERE first = ? AND second = ? " +
		"ORDER BY id DESC PER PARTITION LIMIT 1"
	stmt_list_conversation_after = "SELECT id, status, message FROM message.conversation " +
		"WHERE first = ? AND second = ? AND id > ?" +
		"ORDER BY id ASC PER PARTITION LIMIT ?"
	stmt_send_message = "INSERT INTO message.conversation " +
		"(first, second, id, status, message) VALUES (?, ?, ?, ?, ?)"
)

func scanMessages(scanner gocql.Scanner, user int64, friend int64) []*rpc.Message {
	messages := make([]*rpc.Message, 0)
	for scanner.Next() {
		var (
			id      int64
			status  byte
			message string
		)
		if err := scanner.Scan(&id, &status, &message); err != nil {
			continue
		}
		msg := &rpc.Message{
			Id:      id,
			Content: message,
		}
		if (status & 1) == 0 {
			msg.FromUserId, msg.ToUserId = user, friend
		} else {
			msg.FromUserId, msg.ToUserId = friend, user
		}
		messages = append(messages, msg)
	}
	return messages
}

func ListMessages(user int64, friend int64, after int64, limit int) []*rpc.Message {
	if user > friend {
		user, friend = friend, user
	}

	scanner := session.Query(
		stmt_list_conversation_after,
		user,
		friend,
		snowy.FromUpperTime(time.UnixMilli(after)),
		limit,
	).Iter().Scanner()

	messages := scanMessages(scanner, user, friend)

	return reversed(messages)
}

func reversed(messages []*rpc.Message) []*rpc.Message {
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}

func Send(user int64, friend int64, message string) utils.ErrorCode {
	var status byte
	if user > friend {
		status = 1
		user, friend = friend, user
	}
	if err := session.Query(
		stmt_send_message,
		user,
		friend,
		snowy.ID(),
		status,
		message,
	).Exec(); err != nil {
		return utils.ErrorDatabaseError
	}

	return utils.ErrorOk
}

func LatestMessages(ctx context.Context, user int64, friends []int64) []*rpc.Message {
	limit := semaphore.NewWeighted(8)
	m := sync.Mutex{}
	messages := make([]*rpc.Message, 0)
	for _, friend := range friends {
		final_friend := friend
		if limit.Acquire(ctx, 1) != nil {
			break
		}
		go func() {
			var a, b int64
			if user > final_friend {
				a, b = final_friend, user
			} else {
				a, b = user, final_friend
			}
			scanner := session.Query(
				stmt_list_latest_conversation,
				a,
				b,
			).Iter().Scanner()
			latest := scanMessages(scanner, a, b)
			if len(latest) > 0 {
				m.Lock()
				if ctx.Err() == nil {
					messages = append(messages, latest[0])
				}
				m.Unlock()
			}

			limit.Release(1)
		}()
	}
	// Wait for goroutines to finish writing.
	limit.Acquire(context.Background(), 8)
	return messages
}
