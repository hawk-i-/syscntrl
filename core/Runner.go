package core

import "github.com/jinzhu/gorm"

// Runner object
type Runner struct {
	gorm.Model
	Name       string `gorm:"not null;unique"`
	ParentTask uint
}

// AutoMigrate as
func (r Runner) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&r)
}
