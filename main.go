package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/{key}", KeyValuePutHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8181", router))
}



var ErrorNoSuchKey = errors.New("no such key")

var store = make(map[string]string)

func Store() map[string]string {
	return store
}

func Put(key, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := Store()[key]
	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}

// v1/key/{key}
func KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = Put(key, string(value))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}