package core

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Context struct {
	Config     *viper.Viper
	DBProvider func() (*gorm.DB, error)
}

var context Context
var initialized bool

func InitializePackage(c Context) (err error) {
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

func migrateSchema() (err error) {
	db, err := context.DBProvider()

	if err != nil {
		return
	}

	// db.Model(&Task{}).Related(&Runner{})
	db.AutoMigrate(&Runner{}, &Task{})

	seedData(db)
	return
}

func seedData(db *gorm.DB) {
	task := Task{
		Description: "This is sample task",
		Name:        "SAMPLE_TASK",
		Runners: []Runner{
			{
				Name:        "SAMPLE_RUNNER",
				Description: "This is sample description",
			},
		},
		SubTasks: []Task{
			{
				Name:        "SAMPLE_SUBTASK",
				Description: "This is sample subtask",
			},
		},
	}

	db.Create(&task)
}
