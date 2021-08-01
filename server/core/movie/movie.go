package coremovie

import "gorm.io/gorm"

type service struct {
	pg *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		pg: db,
	}
}

type Service interface {
	ListMovie()
}

func (s *service) ListMovie() {

}
