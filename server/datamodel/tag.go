package datamodel

import (
	"time"
)

type Tag struct {
	ID uint `gorm:"primaryKey"`
	// movies    []Movie
	// users     []User
	tag       string
	timestamp time.Time
}
