package cql_test

import (
	"common/snowy"
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
	assert.Nil(t, err)
	if len(inbox) != 0 {
		assert.True(t, time.Now().After(snowy.Time(inbox[0])))
	}

	assert.Nil(t, cql.PushInboxes(int64(item1), []int64{1}))
	assert.Nil(t, cql.PushInboxes(int64(item2), []int64{1}))

	inbox, err = cql.ListInbox(1, time.Now(), 30)
	assert.Nil(t, err)
	if assert.GreaterOrEqual(t, len(inbox), 2) {
		assert.Equal(t, int64(item2), inbox[0])
		assert.Equal(t, int64(item1), inbox[1])
	}
}
