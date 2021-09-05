package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Actor struct {
	ActorID   uuid.UUID `gorm:"primaryKey" json:"actorId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Movies    []Movie   `gorm:"many2many:movie_actors;" json:"movies"`
}

func (a *Actor) BeforeCreate(tx *gorm.DB) (err error) {
	a.ActorID = uuid.New()
	return
}
