package main

// go get -u github.com/gorilla/mux
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type tag struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	// init router
	router := mux.NewRouter()

	// route handler / endpoints
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/data", getData).Methods("GET")

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

// ==================================
// Home Page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "./../mysql.html")
}

// ==================================
// Get all data
func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataku []tag

	// mysql connection
	db, err := sql.Open("mysql", "lintang:12345@tcp(localhost:3306)/mqttjs")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM mqttjs")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var tag tag
		err := results.Scan(&tag.ID, &tag.Message, &tag.Time)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// fmt.Println(tag)
		fmt.Println(tag.ID, tag.Message, tag.Time)
		dataku = append(dataku, tag)
	}
	json.NewEncoder(w).Encode(dataku)
}
