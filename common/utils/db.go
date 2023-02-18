package utils

import (
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gocql/gocql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func GormDialector() gorm.Dialector {
	dsn := Env.Rdbms
	if strings.HasPrefix(dsn, "file::memory") {
		return mysql.Open("mdouyin:test@tcp(127.0.0.1:3306)/mdouyin?charset=utf8mb4&parseTime=True&loc=Local")
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

func GetSession(cluster *gocql.ClusterConfig) (session *gocql.Session, err error) {
	duration := time.Second * 10
	session, err = cluster.CreateSession()
	for i := 0; i < 5 && err != nil; i++ {
		klog.Warnf("sleep for %s until next try: %s", duration.String(), err.Error())
		time.Sleep(duration)
		duration = duration * 2
		session, err = cluster.CreateSession()
	}
	return
}
