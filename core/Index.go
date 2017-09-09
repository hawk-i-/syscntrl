package core

import (
	"github.com/spf13/viper"
	"errors"
	"github.com/jinzhu/gorm"
)

type Context struct {
	Config *viper.Viper
	DBProvider func() (*gorm.DB,error)
}

var context Context
var initialized bool

func InitializePackage (c Context) (err error) {
	if initialized {
		err = errors.New("package is already initialized")
		return
	}
	context = c
	err = migrateSchema()

	if err != nil {
		return
	}

	initialized = true
	return
}

func migrateSchema () (err error) {
	db, err := context.DBProvider()

	if err != nil {
		return
	}

	Runner{}.AutoMigrate(db)
	Task{}.AutoMigrate(db)

	return
}