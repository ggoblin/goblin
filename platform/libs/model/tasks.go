package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskId    string `gorm:"unique_index"`
	Name      string
	Point     float64
	StartDate time.Time
	EndDate   time.Time
	Raw       string `sql:"json"`
}
