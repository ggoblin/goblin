package model

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Id          string
	Name        string
	Email       string
	DisplayName string
}
