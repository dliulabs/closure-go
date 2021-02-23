package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type City struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity City
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		log.Printf("Got METHOD: %s, not allowed`\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

// Pre-Request Middleware
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type, please send application/json"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Printf("Cookie: %s, Value: %v", cookie.Name, cookie.Value)
		log.Println("Currently in the set server time middleware")
	})
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	// http.Handle("/city", filterContentType(mainLogicHandler))
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))
	log.Println(`curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}'`)
	log.Println(`curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'`)
	log.Println(`curl -H "Content-Type: application/json" -X PATCH http://localhost:8000/city -d '{"name":"Boston", "area":89}'`)
	http.ListenAndServe(":8000", nil)
}
