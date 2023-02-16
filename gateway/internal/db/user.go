package db

import (
	"common/snowy"
	"common/utils"

	"github.com/alexedwards/argon2id"
)

const (
	UserFieldId       = "ID"
	UserFieldName     = "Name"
	UserFieldNickname = "NickName"
	UserFieldPassword = "PasswordHash"
	MinPasswordLength = 6
)

type UserDO struct {
	Id           int64  `gorm:"<-:create;primaryKey;autoIncrement=false"`
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
		Id:           snowy.ID(),
		Name:         name,
		Nickname:     name,
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

func FindUserById(id int64, fields ...string) (*UserDO, error) {
	user := UserDO{Id: id}
	return FindUserWith(&user, fields...)
}

func FindUsersByIds(ids []int64) (users []UserDO, err error) {
	if len(ids) == 0 {
		return []UserDO{}, nil
	}

	if err = db.Limit(len(ids)).Find(&users, ids).Error; err != nil {
		db.Error = nil
		return
	}
	return
}

func UserExists(id int64) bool {
	return db.Limit(1).Find(&UserDO{Id: id}).RowsAffected == 1
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
