package coremovie

import (
	"fmt"

	"github.com/Ekkawin/golang/server/datamodel"
	"gorm.io/gorm"
)

type service struct {
	pg *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		pg: db,
	}
}

type Service interface {
	ListMovie() datamodel.Movie
}

func (s *service) ListMovie() datamodel.Movie {
	var movie datamodel.Movie
	fmt.Println("hello%+v", movie)
	s.pg.First(&movie)
	fmt.Println("hello2%+v", movie)
	return movie

}
