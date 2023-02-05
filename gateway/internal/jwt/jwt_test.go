package jwt_test

import (
	"encoding/hex"
	"gateway/internal/jwt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	assert.NotNil(t, jwt.Init("_", 0))

	assert.NotNil(t, jwt.Init(hex.EncodeToString([]byte("secret")), 0))

	duration := time.Second * 2
	assert.Nil(t, jwt.Init(hex.EncodeToString([]byte("secret")), duration))

	token, err := jwt.NewAuthorization(1234567890, "the_username")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	id, name, err := jwt.Validate(token)
	assert.Nil(t, err)
	assert.Equal(t, int64(1234567890), id)
	assert.Equal(t, "the_username", name)

	time.Sleep(duration)
	id, name, err = jwt.Validate(token)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), id)
	assert.Equal(t, "", name)
}
