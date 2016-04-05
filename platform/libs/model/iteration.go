package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type JsonTime struct {
	time.Time
}

const ctLayout = "2006/01/02|15:04:05"

func (ct *JsonTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(ctLayout, string(b))
	return err
}

func (ct *JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *JsonTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Iteration struct {
	gorm.Model
	IterationId *string    `gorm:"unique_index"`
	Name        *string    `gorm:"not null"`
	StartDate   JsonTime   `gorm:"-" json:"-"`
	EndDate     JsonTime   `gorm:"-" json:"-"`
	StartTime   *time.Time `gorm:"not null"`
	EndTime     *time.Time `gorm:"not null"`
}

func (it *Iteration) SetTime() {
	it.StartTime = &it.StartDate.Time
	it.EndTime = &it.EndDate.Time
}
