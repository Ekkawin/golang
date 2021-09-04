package coremovie

import (
	"fmt"

	"github.com/Ekkawin/golang/server/datamodel"
	"github.com/google/uuid"
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
	ListMovie() []datamodel.Movie
	CreateMovie(movie datamodel.Movie) error
	UpdateMovie(movie datamodel.Movie) (uuid.UUID, error)
}

func (s *service) ListMovie() []datamodel.Movie {
	var movies []datamodel.Movie

	db := s.pg.Preload("Director")
	db.Preload("Actors").Preload("Generes").Limit(10).Offset(1).Order("Title desc").Find(&movies)
	return movies

}

func (s *service) CreateMovie(movie datamodel.Movie) error {
	result := s.pg.Create(&movie)
	fmt.Println(result, "result")
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (s *service) UpdateMovie(movie datamodel.Movie) (uuid.UUID, error) {

	result := s.pg.Save(&movie)
	fmt.Println(result, "result")
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return movie.MovieID, nil

}
