package apiuploadfile

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadFileHandler struct {
	db *gorm.DB
}

func newUploadFileHandler(db *gorm.DB) UploadFileHandlerController {
	return &UploadFileHandler{
		db: db,
	}
}

type UploadFileHandlerController interface {
	CreateMovieGenere(c *gin.Context)
}
