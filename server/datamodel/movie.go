package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	MovieID     uuid.UUID `gorm:"primaryKey"`
	Title       string
	Description string
	Actor       []Actor
	Generes     []Genere
	Director    Director
	Rank        string
	Year        string
	RunTime     string
	Vote        string
	Rating      string
	Revenue     string
	MetaScore   string
}
type Actor struct {
	ActorID   uuid.UUID `gorm:"primaryKey"`
	FirstName string
	LastName  string
}
type Director struct {
	DirectorID uuid.UUID `gorm:"primaryKey"`
	FirstName  string
	LastName   string
}

type Genere struct {
	GenereID uuid.UUID `gorm:"primaryKey"`
	Name     string
}

func (g *Genere) BeforeCreate(tx *gorm.DB) (err error) {
	g.GenereID = uuid.New()
	return
}

func (d *Director) BeforeCreate(tx *gorm.DB) (err error) {
	d.DirectorID = uuid.New()
	return
}
func (a *Actor) BeforeCreate(tx *gorm.DB) (err error) {
	a.ActorID = uuid.New()
	return
}
