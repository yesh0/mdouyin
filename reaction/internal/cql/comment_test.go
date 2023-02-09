package cql_test

import (
	"common/utils"
	"log"
	"math/rand"
	"reaction/internal/cql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := cql.Init("127.0.0.1"); err != nil {
		log.Fatalln(err)
	}

	m.Run()
}

func TestCommentList(t *testing.T) {
	rand.Seed(time.Now().Unix())
	id := rand.Int63()
	list, err := cql.ListComment(id, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Empty(t, list)

	ids := make([]int64, 0, 25)
	for i := 0; i < 25; i++ {
		comment, err := cql.AddComment(id, 1, "comment")
		assert.Equal(t, utils.ErrorOk, err)
		assert.NotNil(t, comment)
		assert.Equal(t, "comment", comment.Content)
		assert.Equal(t, int64(1), comment.User.Id)
		assert.Equal(t, time.Now().Format("01-02"), comment.CreateDate)

		ids = append(ids, comment.Id)
	}

	list, err = cql.ListComment(id, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 25)

	for _, comment := range ids[0:15] {
		assert.Equal(t, utils.ErrorOk, cql.DeleteComment(id, comment, 1))
	}
	list, err = cql.ListComment(id, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 10)

	for _, comment := range ids[15:] {
		assert.Equal(t, utils.ErrorOk, cql.DeleteComment(id, comment, 2))
	}
	list, err = cql.ListComment(id, time.Now(), 30)
	assert.Equal(t, utils.ErrorOk, err)
	assert.Len(t, list, 10)
}
