package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	dsn := "host=localhost user=aek password=aek dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(db, "db")
	fmt.Println(err, "err")

	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// myRouter := mux.NewRouter()
	myRouter := gin.Default()
	myRouter.Use(cors.Default())

	// myRouter.GET("/aek", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "hi my name is aek")
	// })
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "aek")
	// })
	myRouter.GET("/:userName/:password", func(c *gin.Context) {
		userName := c.Param("userName")
		password := c.Param("password")
		fmt.Printf("userName: %v; password: v%", userName, password)
		status := c.Writer.Status()
		fmt.Println("status", status)
		c.JSON(200, gin.H{
			"status":   "success",
			"userName": userName,
			"password": password,
		})

	})

	myRouter.Run(":8080")
	// http.ListenAndServe(":8080", myRouter)
}
