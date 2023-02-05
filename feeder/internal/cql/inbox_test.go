package cql_test

import (
	"common/snowy"
	"common/utils"
	"feeder/internal/cql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInbox(t *testing.T) {
	item1 := snowy.ID()
	item2 := snowy.ID()
	time.Sleep(5 * time.Millisecond)

	inbox, err := cql.ListInbox(1, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	if len(inbox) != 0 {
		assert.True(t, time.Now().After(snowy.Time(inbox[0])))
	}

	assert.Equal(t, utils.ErrorOk, cql.PushInboxes(int64(item1), []int64{1}))
	assert.Equal(t, utils.ErrorOk, cql.PushInboxes(int64(item2), []int64{1}))

	inbox, err = cql.ListInbox(1, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	if assert.GreaterOrEqual(t, len(inbox), 2) {
		assert.Equal(t, int64(item2), inbox[0])
		assert.Equal(t, int64(item1), inbox[1])
	}
}
