package datamodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Director struct {
	DirectorID uuid.UUID `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Movies     []Movie `gorm:"foreignKey:DirectorID;"`
}

func (d *Director) BeforeCreate(tx *gorm.DB) (err error) {
	d.DirectorID = uuid.New()
	return
}
