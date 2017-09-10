package main

import (
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB_DIALECT = "mysql"

func getDBGenerator (dbConfig *viper.Viper) (func() (*gorm.DB,error)) {
	return func() (db *gorm.DB, err error) {
		db, err = gorm.Open(DB_DIALECT, dbConfig.Sub(DB_DIALECT).GetString("path"))
		return
	}
}