package main

import (
	"fmt"
	"net/http"

	coremovie "github.com/Ekkawin/golang/server/core/movie"
	// initializeDb "github.com/Ekkawin/golang/server/datamodel/initialDB"
	"github.com/Ekkawin/golang/server/pkg/apimovie"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	aek string
	abc string
	a23 string
}

func main() {

	dsn := "host=localhost user=aek password=aek dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// initializeDb.InitializeInformationDb(db)
	fmt.Println(err)

	// // myRouter := mux.NewRouter()
	myRouter := gin.Default()
	myRouter.Use(cors.Default())

	movieService := coremovie.NewService(db)
	movieController := apimovie.MovieHandler(movieService)

	myRouter.GET("/movies", movieController.ListMovieController)

	// au := apiuser.NewGetUserHandler(db)

	// // myRouter.GET("/aek", func(w http.ResponseWriter, r *http.Request) {
	// // 	fmt.Fprintf(w, "hi my name is aek")
	// // })
	// // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	fmt.Fprintf(w, "aek")
	// // })
	// myRouter.GET("/users", au.ListHandler)
	// myRouter.POST("/user", au.CreateUser)
	// myRouter.POST("/upload-file")

	myRouter.Run(":8080")
	http.ListenAndServe(":8080", myRouter)
}
