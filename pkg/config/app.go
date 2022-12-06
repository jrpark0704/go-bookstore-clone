package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:sde585766!@/simplerest?charset=utf8&&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("db not connected")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
