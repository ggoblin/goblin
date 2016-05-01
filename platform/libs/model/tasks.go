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
	StartDate   *JsonTime  `gorm:"-"`
	EndDate     *JsonTime  `gorm:"-"`
	StartTime   *time.Time `gorm:"not null" json:"-"`
	EndTime     *time.Time `gorm:"not null" json:"-"`
	Raw         *string    `sql:"json"`
}

func (it *Task) SetDate() {
	it.StartDate = &JsonTime{it.StartTime}
	it.EndDate = &JsonTime{it.EndTime}
}

func (it *Task) SetTime() {
	it.StartTime = it.StartDate.Time
	it.EndTime = it.EndDate.Time
}
