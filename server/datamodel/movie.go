package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	MovieID     uuid.UUID `gorm:"primaryKey"`
	Title       string
	Description string
	Actors      []Actor  `gorm:"many2many:movie_actors; OnDelete:SET NULL;"`
	Generes     []Genere `gorm:"many2many:movie_generes; OnDelete:SET NULL;;"`
	DirectorID  uuid.UUID
	Director    Director
	Rank        string
	Year        string
	RunTime     string
	Vote        string
	Rating      string
	Revenue     string
	MetaScore   string
}

func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	m.MovieID = uuid.New()
	return
}
