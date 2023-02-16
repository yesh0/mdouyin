package utils

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GormDialector() gorm.Dialector {
	dsn := Env.Rdbms
	if strings.HasPrefix(dsn, "file::memory") {
		return sqlite.Open(dsn)
	} else {
		return mysql.Open(dsn)
	}
}
