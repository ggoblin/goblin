package model

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	MemberId    string `gorm:"unique_index"`
	Name        string
	Email       string
	DisplayName string
}
