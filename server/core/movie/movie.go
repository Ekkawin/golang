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
	GetMovie(movieId uuid.UUID) (datamodel.Movie, error)
	CreateMovie(movie datamodel.Movie) error
	UpdateMovie(movie datamodel.Movie) (uuid.UUID, error)
	DeleteMovie(movieId uuid.UUID) error
}

func (s *service) ListMovie() []datamodel.Movie {
	var movies []datamodel.Movie

	db := s.pg.Preload("Director")
	db.Preload("Actors").Preload("Generes").Limit(10).Offset(1).Order("Title desc").Find(&movies)
	return movies

}

func (s *service) CreateMovie(movie datamodel.Movie) error {
	result := s.pg.Create(&movie)

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
func (s *service) DeleteMovie(movieId uuid.UUID) error {
	// actors := movie.Actors
	// deleteActor := s.pg.Delete(&actors)
	// if deleteActor.Error != nil {
	// 	return deleteActor.Error
	// }
	// generes := movie.Generes
	// deleteGenere := s.pg.Delete(&generes)

	// if deleteGenere.Error != nil {
	// 	return deleteGenere.Error

	// }
	// director := movie.Director
	// deleteDirector := s.pg.Delete(&director)

	// if deleteDirector.Error != nil {
	// 	return deleteDirector.Error
	// var movie datamodel.Movie
	// s.pg.Where("movie_id = ?", movieId).Preload("Actors").Preload("Generes").First(&movie)
	// fmt.Println(movie, "movie")
	// result := s.pg.Unscoped().Delete(&movie)

	// if result.Error != nil {
	// 	return result.Error
	// }
	return nil
	// return movie
}

func (s *service) GetMovie(movieId uuid.UUID) (datamodel.Movie, error) {
	var movie datamodel.Movie
	result := s.pg.Where("movie_id = ?", movieId).Preload("Actors").Preload("Generes").First(&movie)
	if result.Error != nil {
		return datamodel.Movie{}, result.Error

	}
	return movie, nil

}
