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
	ListMovie() *gorm.DB
}

func (s *service) ListMovie() *gorm.DB {
	var movie datamodel.Movie
	fmt.Println("hello")
	result := s.pg.First(&movie)
	fmt.Println("hello2")
	return result

}
