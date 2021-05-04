package apiuser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func Initialize(db *gorm.DB) Handler {
	return &UserHandler{
		db: db,
	}
}

type Handler interface {
	CreateUser(c *gin.Context)
}
