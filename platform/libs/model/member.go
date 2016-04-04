package model

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	MemberId    string `gorm:"unique_index"`
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	DisplayName string `gorm:"not null"`
}
