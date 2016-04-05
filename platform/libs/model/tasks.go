package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskId    *string    `gorm:"unique_index"`
	Name      *string    `gorm:"not null"`
	Point     *float64   `gorm:"not null"`
	StartDate JsonTime   `gorm:"-"`
	EndDate   JsonTime   `gorm:"-"`
	StartTime *time.Time `gorm:"not null" json:"-"`
	EndTime   *time.Time `gorm:"not null" json:"-"`
	Raw       *string    `sql:"json"`
}

func (it *Task) SetTime() {
	if it.StartTime == nil {
		it.StartTime = &it.StartDate.Time
		it.EndTime = &it.EndDate.Time

	} else {
		it.StartDate.Time = *it.StartTime
		it.EndDate.Time = *it.EndTime
	}

}
