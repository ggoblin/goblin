package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Iteration struct {
	gorm.Model
	IterationId string    `gorm:"unique_index"`
	Name        string    `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
}
