package main

import (
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB_DIALECT = "postgres"

func getDBGenerator (dbConfig *viper.Viper) (func() (*gorm.DB,error)) {
	return func() (db *gorm.DB, err error) {
		db, err = gorm.Open(DB_DIALECT, dbConfig.GetString("path"))
		return
	}
}