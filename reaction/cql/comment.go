package cql

import (
	"common/kitex_gen/douyin/rpc"
	"common/snowy"
	"common/utils"
	"time"
)

const (
	stmt_list_comments = "SELECT item, author, content, removed FROM reaction.comment " +
		"WHERE video = ? AND item < ? " +
		"ORDER BY item DESC PER PARTITION LIMIT ?"
	stmt_insert_comment = "INSERT INTO reaction.comment " +
		"(video, item, author, content, removed) VALUES (?, ?, ?, ?, 0)"
	stmt_delete_comment = "UPDATE reaction.comment SET removed = 1 " +
		"WHERE video = ? AND item = ? AND author = ?"
)

func ListComment(video int64, latest time.Time, limit int) ([]*rpc.Comment, utils.ErrorCode) {
	scanner := session.Query(
		stmt_list_comments,
		video,
		snowy.FromLowerTime(latest),
		limit,
	).Iter().Scanner()

	items := make([]*rpc.Comment, 0, 16)
	for scanner.Next() {
		var (
			item    int64
			author  int64
			content string
			removed int8
		)
		if err := scanner.Scan(&item, &author, &content, &removed); err != nil {
			return nil, utils.ErrorDatabaseError
		}
		if removed != 0 {
			continue
		}
		items = append(items, comment(item, author, content))
	}
	return items, utils.ErrorOk
}

func comment(id int64, user int64, content string) *rpc.Comment {
	return &rpc.Comment{
		Id:         id,
		User:       &rpc.User{Id: user},
		Content:    content,
		CreateDate: snowy.Time(id).Format("01-02"),
	}
}

func AddComment(video int64, user int64, content string) (*rpc.Comment, utils.ErrorCode) {
	id := snowy.ID()
	if err := session.Query(
		stmt_insert_comment,
		video,
		id,
		user,
		content,
	).Exec(); err != nil {
		return nil, utils.ErrorDatabaseError
	}
	return comment(id, user, content), utils.ErrorOk
}

func DeleteComment(video int64, id int64, user int64) utils.ErrorCode {
	if err := session.Query(stmt_delete_comment, video, id, user).Exec(); err != nil {
		return utils.ErrorDatabaseError
	}
	return utils.ErrorOk
}
