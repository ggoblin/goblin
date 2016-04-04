package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskId    string    `gorm:"unique_index"`
	Name      string    `gorm:"not null"`
	Point     float64   `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Raw       string    `sql:"json"`
}
