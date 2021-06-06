package apiuser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewGetUserHandler(db *gorm.DB) UserHandlerController {
	return &UserHandler{
		db: db,
	}
}

type UserHandlerController interface {
	CreateUser(c *gin.Context)
}
