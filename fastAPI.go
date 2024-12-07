package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Starting Gorilla Mux usage!! \nGo to localhost:8080/hello for routing in action"))
	}).Methods("GET")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
