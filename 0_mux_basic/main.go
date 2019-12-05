package main

// go get -u github.com/gorilla/mux
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

// ==================================
// Home Page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "./../index.html")
}

// ==================================
// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// ==================================
// Get a book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get param

	// loop through books & find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&book{})
}

// ==================================
// create book POST
// {
// 	"isbn": "99999",
// 	"title": "Buku Baru",
// 	"author": {
// 		"firstname": "Lintang",
// 		"lastname": "Wisesa"
// 	}
// }

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // create random ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// ==================================
// update book PUT
// {
// 	"isbn": "99999",
// 	"title": "Buku Baru",
// 	"author": {
// 		"firstname": "Lintang",
// 		"lastname": "Wisesa"
// 	}
// }
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// ==================================
// delete book DELETE
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init router
	router := mux.NewRouter()

	// data
	books = append(books, book{ID: "1", Isbn: "123456", Title: "Buku Satu", Author: &author{Firstname: "Andi", Lastname: "Susilo"}})
	books = append(books, book{ID: "2", Isbn: "123457", Title: "Buku Dua", Author: &author{Firstname: "Budi", Lastname: "Susilo"}})
	books = append(books, book{ID: "3", Isbn: "123458", Title: "Buku Tiga", Author: &author{Firstname: "Caca", Lastname: "Susilo"}})
	books = append(books, book{ID: "4", Isbn: "123459", Title: "Buku Empat", Author: &author{Firstname: "Deni", Lastname: "Susilo"}})
	books = append(books, book{ID: "5", Isbn: "123460", Title: "Buku Lima", Author: &author{Firstname: "Fafa", Lastname: "Susilo"}})

	// route handler / endpoints
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run server basic
	// log.Fatal(http.ListenAndServe(":1234", router))

	// run server with config
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:1234",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Connected to port 1234")
	log.Fatal(srv.ListenAndServe())
}
