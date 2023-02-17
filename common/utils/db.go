package utils

import (
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func GormDialector() gorm.Dialector {
	dsn := Env.Rdbms
	if strings.HasPrefix(dsn, "file::memory") {
		return sqlite.Open(dsn)
	} else {
		return mysql.Open(dsn)
	}
}

func Open() (*gorm.DB, error) {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()
	db, err := gorm.Open(GormDialector(), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	if sqlDB, err := db.DB(); err != nil {
		return nil, err
	} else {
		sqlDB.SetMaxOpenConns(10)
	}

	return db, nil
}
