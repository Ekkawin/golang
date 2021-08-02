package apimovie

import (
	"fmt"

	coremovie "github.com/Ekkawin/golang/server/core/movie"
	"github.com/gin-gonic/gin"
)

type moviecontroller struct {
	service coremovie.Service
}

func MovieHandler(router *gin.Engine, service coremovie.Service) {
	mc := movieController(router, service)
	router.GET("/movies", mc.ListMovieController)
}

func movieController(router *gin.Engine, service coremovie.Service) MovieController {
	return moviecontroller{
		service: service,
	}
}

type MovieController interface {
	ListMovieController(c *gin.Context)
}

func (mc moviecontroller) ListMovieController(c *gin.Context) {
	fmt.Println("hi")
	a := mc.service.ListMovie()
	fmt.Println(a, "a")

}
