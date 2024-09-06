package db

import (
	"github.com/e2site/sharks-go-lib/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dsn string) {
	inst, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.CheckEndLogFatal(err)
	db = inst
}

func GetDB() *gorm.DB {
	return db
}
