package core

import "github.com/jinzhu/gorm"

// Trigger Object
type Trigger struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Description  string
	ParentTaskID uint
}
