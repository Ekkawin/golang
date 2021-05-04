package main

import (
	"fmt"

	"github.com/Ekkawin/golang/server/datamodel"
	"github.com/Ekkawin/golang/server/pkg/apiuser"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=aek password=aek dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(db, "db")
	fmt.Println(err, "err")

	db.AutoMigrate(&datamodel.Product{})
	db.AutoMigrate(&datamodel.User{})

	// Create
	// db.Create(&datamodel.Product{Code: "D43", Price: 10})

	// myRouter := mux.NewRouter()
	myRouter := gin.Default()
	myRouter.Use(cors.Default())
	au := apiuser.Initialize(db)

	// myRouter.GET("/aek", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "hi my name is aek")
	// })
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "aek")
	// })
	myRouter.POST("/user", au.CreateUser)

	myRouter.Run(":8080")
	// http.ListenAndServe(":8080", myRouter)
}
