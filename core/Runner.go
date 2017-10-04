package core

import "github.com/jinzhu/gorm"

// Runner object
type Runner struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Description  string
	ParentTaskID uint
}
