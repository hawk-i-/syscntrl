package core

import "github.com/jinzhu/gorm"

// Task object
type Task struct {
	gorm.Model
	Name         string `gorm:"not null;unique_index"`
	Description  string
	Runners      []Runner `gorm:"ForeignKey:ParentTaskID"`
	SubTasks     []Task   `gorm:"ForeignKey:ParentTaskID"`
	ParentTaskID uint
}
