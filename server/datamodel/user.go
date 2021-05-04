package datamodel

type User struct {
	ID uint `gorm:"primaryKey type:uuid;default:uuid_generate_v4()"`

	Email    string `json:"email"`
	Password string `json:"password"`
}
