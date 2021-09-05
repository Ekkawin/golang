package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genere struct {
	GenereID uuid.UUID `gorm:"primaryKey" json:"genereId"`
	Name     string    `json:"name"`
	Movies   []Movie   `gorm:"many2many:movie_generes;" json:"movies"`
}

func (g *Genere) BeforeCreate(tx *gorm.DB) (err error) {
	g.GenereID = uuid.New()
	return
}
