package db_test

import (
	"gateway/db"
	"testing"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/stretchr/testify/assert"
)

const test_password = "test_password_1234567891011121314151617181920"
const test_username = "new_name"

func TestArgon2idHash(t *testing.T) {
	hash, err := argon2id.CreateHash(
		test_password,
		argon2id.DefaultParams,
	)
	assert.Nil(t, err)
	user := db.UserDO{
		PasswordHash: hash,
	}
	assert.Nil(t, user.VerifyPassword(test_password))
}

func TestUserCreation(t *testing.T) {
	user, err := db.CreateUser(test_username, test_password)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Nil(t, user.VerifyPassword(test_password))

	assert.NotEqual(t, 0, user.Id)

	user, err = db.FindUserByName(test_username)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Nil(t, user.VerifyPassword(test_password))

	user, err = db.FindUserById(user.Id, "Name", "CreatedAt")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.VerifyPassword(test_password))
	assert.Less(t, time.Until(user.CreatedAt).Abs().Minutes(), float64(1))
}
