package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Director struct {
	DirectorID uuid.UUID `gorm:"primaryKey" json:"directorId"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Movies     []Movie   `gorm:"foreignKey:DirectorID;" json:"movies"`
}

func (d *Director) BeforeCreate(tx *gorm.DB) (err error) {
	d.DirectorID = uuid.New()
	return
}
