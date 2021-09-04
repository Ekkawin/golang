package apimovie

import (
	"fmt"
	"net/http"

	coremovie "github.com/Ekkawin/golang/server/core/movie"
	"github.com/Ekkawin/golang/server/datamodel"
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
	CreateMovieController(c *gin.Context)
	UpdateMovieController(c *gin.Context)
}

func (mc *movieListHandler) ListMovieController(c *gin.Context) {
	fmt.Println("hi")
	a := mc.service.ListMovie()
	c.JSON(http.StatusOK, a)

}
func (mc *movieListHandler) CreateMovieController(c *gin.Context) {
	var movie datamodel.Movie
	if err := c.BindJSON(&movie); err != nil {
		fmt.Println(("err"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mc.service.CreateMovie(movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}
func (mc *movieListHandler) UpdateMovieController(c *gin.Context) {
	var movie datamodel.Movie
	id := c.PostForm("MovieID")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "movieId is required"})
		return
	}

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := mc.service.UpdateMovie(movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"movieId": m,
	})

}
