package coreuser

import (
	"fmt"

	"github.com/Ekkawin/golang/server/datamodel"
)

func (s *service) UserCreate(user *datamodel.User) {
	fmt.Printf("userin core: %v", user)

	s.pg.Create(user)

}
