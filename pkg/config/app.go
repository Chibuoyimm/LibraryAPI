package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // the documentation said to do that
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "sql6519166:Zi5crGsTws@/sql6519166?charset=utf8&parseTime=True&loc=Local") // setting up mysql connection "username:password@/databasename"
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
