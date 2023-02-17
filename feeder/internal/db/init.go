package db

import (
	"common/utils"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	if db != nil {
		return fmt.Errorf("db already initialized")
	}

	var err error
	if db, err = utils.Open(); err != nil {
		return err
	}

	if err := db.AutoMigrate(&RelationDO{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&VideoDO{}); err != nil {
		return err
	}

	return nil
}
