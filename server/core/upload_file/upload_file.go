package coreuploadfile

import (
	"github.com/Ekkawin/golang/server/datamodel"
	"gorm.io/gorm"
)

type service struct {
	pg *gorm.DB
}

func NewService(pg *gorm.DB) Service {
	return &service{
		pg: pg,
	}

}

type Service interface {
	CreateMovieGenere(user *datamodel.User)
}
