package cql_test

import (
	"common/kitex_gen/douyin/rpc"
	"common/snowy"
	"common/utils"
	"message/internal/cql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversation(t *testing.T) {
	me := snowy.ID()
	friend := snowy.ID()
	assert.Empty(t, cql.ListMessages(me, friend, 10))

	assert.Equal(t, utils.ErrorOk, cql.Send(me, friend, "Hello"))
	assert.Equal(t, utils.ErrorOk, cql.Send(friend, me, "Hi"))

	check := func(messages []*rpc.Message) {
		assert.Len(t, messages, 2)
		assert.Equal(t, "Hi", messages[0].Content)
		assert.Equal(t, friend, messages[0].FromUserId)
		assert.Equal(t, me, messages[0].ToUserId)
		assert.Nil(t, messages[0].CreateTime)

		assert.Equal(t, "Hello", messages[1].Content)
		assert.Equal(t, me, messages[1].FromUserId)
		assert.Equal(t, friend, messages[1].ToUserId)
		assert.Nil(t, messages[1].CreateTime)
	}
	check(cql.ListMessages(me, friend, 10))
	check(cql.ListMessages(friend, me, 10))

	message := cql.LatestMessages(me, []int64{friend})
	assert.Len(t, message, 1)
	assert.Equal(t, "Hi", message[0].Content)
	message = cql.LatestMessages(friend, []int64{me})
	assert.Len(t, message, 1)
	assert.Equal(t, "Hi", message[0].Content)
}
