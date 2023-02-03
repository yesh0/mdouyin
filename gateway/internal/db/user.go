package db

import (
	"common/utils"

	"github.com/alexedwards/argon2id"
	"github.com/godruoyi/go-snowflake"
	"gorm.io/gorm"
)

const (
	UserFieldId       = "ID"
	UserFieldName     = "Name"
	UserFieldNickname = "NickName"
	UserFieldPassword = "PasswordHash"
	MinPasswordLength = 6
)

type UserDO struct {
	gorm.Model
	Id           uint64 `gorm:"<-:create;primaryKey;autoIncrement=false"`
	Name         string `gorm:"<-:create;uniqueIndex"`
	Nickname     string // Nickname
	PasswordHash string // Password hash & salt & params (Argon2id)
}

func migrateUserTable() error {
	return db.AutoMigrate(&UserDO{})
}

func CreateUser(name string, password string) (*UserDO, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return nil, utils.ErrorInternalError.Wrap(err)
	}

	// TODO: Make this requirement configurable
	if len(password) < MinPasswordLength {
		return nil, utils.ErrorPasswordLength
	}

	user := UserDO{
		Id:           snowflake.ID(),
		Name:         name,
		Nickname:     "New User",
		PasswordHash: hash,
	}

	// TODO: Cache
	if err := db.Create(&user).Error; err != nil {
		return nil, utils.ErrorUsernameConflict
	}

	return &user, nil
}

func FindUserWith(user *UserDO, fields ...string) (*UserDO, error) {
	query := db
	if len(fields) != 0 {
		query = db.Select(fields)
	}
	if err := query.First(&user).Error; err != nil {
		return nil, utils.ErrorNoSuchUser
	}
	return user, nil
}

func FindUserByName(name string, fields ...string) (*UserDO, error) {
	user := UserDO{Name: name}
	return FindUserWith(&user, fields...)
}

func FindUserById(id uint64, fields ...string) (*UserDO, error) {
	user := UserDO{Id: id}
	return FindUserWith(&user, fields...)
}

func (u *UserDO) VerifyPassword(password string) error {
	match, err := argon2id.ComparePasswordAndHash(password, u.PasswordHash)
	if err != nil {
		return utils.ErrorPasswordValidation.Wrap(err)
	}
	if !match {
		return utils.ErrorWrongPassword
	}
	return nil
}
