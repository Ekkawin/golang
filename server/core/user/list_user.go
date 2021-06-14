package coreuser

import (
	"fmt"

	"github.com/Ekkawin/golang/server/datamodel"
)

type Result struct {
	email string
	price uint
}

func (s *service) ListUser() []datamodel.User {
	var users []datamodel.User
	// var result Result

	s.pg.Model(&users).Preload("Products").Find(&users)

	fmt.Printf("userin core: %v", users)
	return users

}
