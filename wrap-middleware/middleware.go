package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		log.Println("The request took", end.Sub(start))
	}
}

func main() {
	http.HandleFunc("/hello", timed(hello))
	log.Println("http://localhost:8000/hello")
	http.ListenAndServe(":8000", nil)
}
