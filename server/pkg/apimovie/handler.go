package apimovie

import (
	"fmt"
	"net/http"

	coremovie "github.com/Ekkawin/golang/server/core/movie"
	"github.com/gin-gonic/gin"
)

type movieListHandler struct {
	service coremovie.Service
}

func MovieHandler(service coremovie.Service) MovieController {
	return &movieListHandler{
		service: service,
	}

}

type MovieController interface {
	ListMovieController(c *gin.Context)
}

func (mc *movieListHandler) ListMovieController(c *gin.Context) {
	fmt.Println("hi")
	a := mc.service.ListMovie()
	c.JSON(http.StatusOK, a)

}
