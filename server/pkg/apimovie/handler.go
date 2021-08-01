package apimovie

import (
	coremovie "github.com/Ekkawin/golang/server/core/movie"
	"github.com/gin-gonic/gin"
)

type moviecontroller struct {
	service coremovie.Service
}

func MovieHandler(router *gin.Engine, service coremovie.Service) {
	mc := movieController(router, service)
	router.GET("/movies", mc.ListMovie)
}

func movieController(router *gin.Engine, service coremovie.Service) MovieController {
	return moviecontroller{
		service: service,
	}
}

type MovieController interface {
	ListMovie(c *gin.Context)
}

func (mc moviecontroller) ListMovie(c *gin.Context) {
	// service := coremovie.ser

}
