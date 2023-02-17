package db

import (
	"common/utils"
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Initializes the global DB instance
func Init() error {
	if db != nil {
		return fmt.Errorf("database already initialized")
	}

	var err error

	db, err = utils.Open()
	if err != nil {
		return err
	}

	err = migrateUserTable()
	if err != nil {
		return err
	}

	return nil
}
