package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Books struct data (model)
type book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *author `json:"author"`
}

// author struct
type author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// data
var books []book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create book
func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// init router
	r := mux.NewRouter()

	// data
	books = append(books, book{ID: "1", Isbn: "123456", Title: "Buku Satu", Author: &author{Firstname: "Andi", Lastname: "Susilo"}})
	books = append(books, book{ID: "2", Isbn: "123457", Title: "Buku Dua", Author: &author{Firstname: "Budi", Lastname: "Susilo"}})
	books = append(books, book{ID: "3", Isbn: "123458", Title: "Buku Tiga", Author: &author{Firstname: "Caca", Lastname: "Susilo"}})
	books = append(books, book{ID: "4", Isbn: "123459", Title: "Buku Empat", Author: &author{Firstname: "Deni", Lastname: "Susilo"}})
	books = append(books, book{ID: "5", Isbn: "123460", Title: "Buku Lima", Author: &author{Firstname: "Fafa", Lastname: "Susilo"}})

	// route handler / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":1234", r))
}
