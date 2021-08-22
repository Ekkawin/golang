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
	ListMovie() []datamodel.Movie
}

func (s *service) ListMovie() []datamodel.Movie {
	var movies []datamodel.Movie
	// var actors []datamodel.Actor

	db := s.pg.Preload("Director")
	db.Preload("Actors").Preload("Generes").Limit(10).Offset(1).Order("Title desc").Find(&movies)

	// Joins("left join directors ON movies.director_id = directors.director_id")

	// .Joins("JOIN directos ON movies.director_id = directors.director_id")

	// s.pg.Model(&actor).Association("Movie").Find(&movies)
	fmt.Println(movies, "movie")

	return movies

}
