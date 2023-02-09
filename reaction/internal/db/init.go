package db

import "gorm.io/gorm"

var db *gorm.DB

func Init(dialector gorm.Dialector) (err error) {
	db, err = gorm.Open(dialector)
	if err != nil {
		return
	}

	if err = db.AutoMigrate(&FavoriteDO{}); err != nil {
		return
	}

	return
}
