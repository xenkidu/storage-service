package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gorilla!\n"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloGoHandler)
	log.Fatal(http.ListenAndServe(":8181", router))
}