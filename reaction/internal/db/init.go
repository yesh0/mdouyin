package db

import (
	"common/utils"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	db, err = utils.Open()
	if err != nil {
		return
	}

	if err = db.AutoMigrate(&FavoriteDO{}); err != nil {
		return
	}

	return
}
