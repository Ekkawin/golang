package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	MovieID     uuid.UUID `gorm:"primaryKey" json:"movieId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Actors      []Actor   `gorm:"many2many:movie_actors; OnDelete:SET NULL;" json:"actors"`
	Generes     []Genere  `gorm:"many2many:movie_generes; OnDelete:SET NULL;" json:"generes"`
	DirectorID  uuid.UUID `json:"directorId"`
	Director    Director  `json:"director"`
	Rank        string    `json:"rank"`
	Year        string    `json:"year"`
	RunTime     string    `json:"runTime"`
	Vote        string    `json:"vote"`
	Rating      string    `json:"rating"`
	Revenue     string    `json:"revenue"`
	MetaScore   string    `json:"metaScore"`
}

func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	m.MovieID = uuid.New()
	return
}
