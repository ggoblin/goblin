package model

import (
	"time"
)

type tasks struct {
	Id        string
	Name      string
	Point     float64
	StartDate time.Time
}
