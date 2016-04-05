package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskId      *string    `gorm:"unique_index"`
	IterationId *string    `gorm:"not null"`
	Name        *string    `gorm:"not null"`
	Point       *float64   `gorm:"not null"`
	StartDate   JsonTime   `gorm:"-" json:"-"`
	EndDate     JsonTime   `gorm:"-" json:"-"`
	StartTime   *time.Time `gorm:"not null"`
	EndTime     *time.Time `gorm:"not null"`
	Raw         *string    `sql:"json"`
}

func (it *Task) SetTime() {
	it.StartTime = &it.StartDate.Time
	it.EndTime = &it.EndDate.Time
}
