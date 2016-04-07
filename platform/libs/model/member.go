package model

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	MemberId    *string `gorm:"unique_index"`
	Name        *string `gorm:"not null"`
	Email       *string `gorm:"not null;unique"`
	DisplayName *string `gorm:"not null"`
}

type IterationMember struct {
	gorm.Model
	MemberId    *string  `gorm:"not null"`
	IterationId *string  `gorm:"not null"`
	MemberFact  *float64 `gorm:"not null"`
}

type TaskMember struct {
	gorm.Model
	MemberId *string `gorm:"not null"`
	TaskId   *string `gorm:"not null"`
}
