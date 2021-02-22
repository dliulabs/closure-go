package main

import (
	"fmt"
	"net/http"
	"log"
)

type Database struct {
	Name string
}

func NewDatabase(name string) Database {
	return Database{name}
}

func main() {
	db := NewDatabase("David")

	http.HandleFunc("/hello", hello(db))
	log.Println("http://localhost:8000/hello")
	http.ListenAndServe(":8000", nil)
}

func hello(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Name)
	}
}
