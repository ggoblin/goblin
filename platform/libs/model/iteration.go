package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type JsonTime struct {
	*time.Time
}

const ctLayout = "2006/01/02|15:04:05"

func (ct *JsonTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	t, err := time.Parse(ctLayout, string(b))
	ct.Time = &t
	return err
}

func (ct *JsonTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + ct.Time.Format(ctLayout) + "\""), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *JsonTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Iteration struct {
	gorm.Model
	IterationId *string    `gorm:"unique_index"`
	Name        *string    `gorm:"not null"`
	StartDate   *JsonTime  `gorm:"-"`
	EndDate     *JsonTime  `gorm:"-"`
	StartTime   *time.Time `gorm:"not null" json:"-"`
	EndTime     *time.Time `gorm:"not null" json:"-"`
}

func (it *Iteration) SetDate() {
	it.StartDate = &JsonTime{it.StartTime}
	it.EndDate = &JsonTime{it.EndTime}
}

func (it *Iteration) SetTime() {
	it.StartTime = it.StartDate.Time
	it.EndTime = it.EndDate.Time
}
