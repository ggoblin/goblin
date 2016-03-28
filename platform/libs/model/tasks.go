package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tasks struct {
	gorm.Model
	Id        string
	Name      string
	Point     float64
	StartDate time.Time
	EndDate   time.Time
	Raw       string `sql:"json"`
}
