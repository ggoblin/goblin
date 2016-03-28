package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Iteration struct {
	gorm.Model
	IterationId string
	Name        string
	StartDate   time.Time
	EndDate     time.Time
}
