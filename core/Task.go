package core

import "github.com/jinzhu/gorm"

// Task object
type Task struct {
	gorm.Model
	Name        string `gorm:"not null;unique_index"`
	Description string
	Runners     []Runner `gorm:"ForeignKey:ParentTask"`
}

// AutoMigrate method
func (t Task) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&t)
}
