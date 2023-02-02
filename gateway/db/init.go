package db

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	db *gorm.DB
)

func InitWithMySql(dsn string) error {
	return Init(mysql.Open(dsn))
}

// Initializes the global DB instance
func Init(dialector gorm.Dialector) error {
	if db != nil {
		return fmt.Errorf("database already initialized")
	}

	var err error

	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return err
	}

	err = migrateUserTable()
	if err != nil {
		return err
	}

	return nil
}
