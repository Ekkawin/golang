package datamodel

type Movie struct {
	ID      uint `gorm:"primaryKey"`
	user    []User
	generes string
	TagID   uint
}
