package db_test

import (
	"common/utils"
	"gateway/internal/db"
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
	user, err := db.FindUserById(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	user, err = db.FindUserByName(test_username)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	user, err = db.CreateUser(test_username, test_password)
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

func TestInsecurePassword(t *testing.T) {
	_, err := db.CreateUser("insecure_user", "12345")
	assert.NotNil(t, err)
	assert.Equal(t, err, utils.ErrorPasswordLength)

	user, err := db.CreateUser("insecure_user", "Ot*f@s_8")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestIds(t *testing.T) {
	u1, err := db.CreateUser("id_user1", "12345678")
	assert.Nil(t, err)
	u2, err := db.CreateUser("id_user2", "12345678")
	assert.Nil(t, err)

	users, err := db.FindUsersByIds([]int64{int64(u1.Id), int64(u2.Id)})
	assert.Nil(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, u1.Name, users[0].Name)
	assert.Equal(t, u2.Name, users[1].Name)
}
