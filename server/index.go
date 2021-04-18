package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/aek", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi my name is aek")
	}).Methods("GET")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "aek")
	})
	http.ListenAndServe(":8080", myRouter)
}
