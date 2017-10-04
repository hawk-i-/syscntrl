package core

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Context struct {
	Config        *viper.Viper
	DBProvider    func() (*gorm.DB, error)
	tokenProvider gTokenProvider
}

var context Context
var initialized bool

func InitializePackage(c Context) (err error) {
	if initialized {
		err = errors.New("package is already initialized")
		return
	}
	context = c
	err = migrateSchema(c)
	context.tokenProvider, err = initTokenProvider(c)
	if err != nil {
		return
	}

	initialized = true
	return
}

func IsInitialized() bool {
	return initialized
}

func migrateSchema(context Context) (err error) {
	db, err := context.DBProvider()

	if err != nil {
		return
	}


	db.AutoMigrate(&Runner{}, &Trigger{}, &Task{})

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

	task2 := Task{
		Description: "This is sample task",
		Name:        "SAMPLE_TASK2",
		Runners: []Runner{
			{
				Name:        "SAMPLE_RUNNER2",
				Description: "This is sample description",
			},
		},
		Triggers: []Trigger{
			{
				Name:        "Test Trigger",
				Description: "This is test trigger",
			},
		},
		SubTasks: []Task{
			task,
		},
	}


	db.Create(&task2)
}
