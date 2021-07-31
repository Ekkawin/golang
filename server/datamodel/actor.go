package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Actor struct {
	ActorID   uuid.UUID `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Movies    []Movie `gorm:"many2many:movie_actors;"`
}

func (a *Actor) BeforeCreate(tx *gorm.DB) (err error) {
	a.ActorID = uuid.New()
	return
}
