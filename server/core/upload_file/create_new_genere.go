package coreuploadfile

import (
	"github.com/Ekkawin/golang/server/datamodel"
)

func (s *service) CreateMovieGenere(user *datamodel.User) {
	// fmt.Printf("userin core: %v", user)

	s.pg.Create(user)

}
