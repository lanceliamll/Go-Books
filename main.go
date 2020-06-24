//golint:ignore
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book presentation
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// store books
var books []Book

// Get Books endpoint
func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Books!")
	log.Println(w, r)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/books", getBooks).Methods("GET")
	http.ListenAndServe(":8080", route)
}
