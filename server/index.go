package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/aek", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi my name is aek")
	}).Methods("GET")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "aek")
	})
	http.ListenAndServe(":8080", myRouter)
}
