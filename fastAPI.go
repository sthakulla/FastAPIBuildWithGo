// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Starting Gorilla Mux usage!! \nGo to localhost:8080/hello for routing in action"))
// 	}).Methods("GET")
// 	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, World!"))
// 	}).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello World!")
// }

//	func main() {
//		http.HandleFunc("/hello", helloHandler)
//		http.ListenAndServe(":8080", nil)
//	}
// package main

// import (
// 	"net/http"

// 	jsoniter "github.com/json-iterator/go"
// )

// var json = jsoniter.ConfigCompatibleWithStandardLibrary

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	response := map[string]string{"message": "Hello, World!"}
// 	json.NewEncoder(w).Encode(response)
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var message string
	err := db.QueryRow("SELECT 'Hello, World!'").Scan(&message)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
