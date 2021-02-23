package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler ...")
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Microsecond)
	w.Write([]byte("OK"))
}

// a closure: returning a func() containing a
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Executing middleware before request phase", start.Format("2006-01-02 15:04:05.000000000"))
		handler.ServeHTTP(w, r)
		end := time.Now()
		fmt.Println("Executing middleware after response phase", end.Format("2006-01-02 15:04:05.000000000"))
	})
}

func main() {
	mainHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainHandler)) // wrapping main logic in middleware
	log.Println("http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
