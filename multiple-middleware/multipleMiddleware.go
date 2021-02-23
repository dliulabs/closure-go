package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	//"strconv"
	//"time"
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

func main() {
	//mainLogicHandler := http.HandlerFunc(mainLogic)
	http.HandleFunc("/city", mainLogic)
	log.Println(`curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}'`)
	log.Println(`curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'`)
	log.Println(`curl -H "Content-Type: application/json" -X PATCH http://localhost:8000/city -d '{"name":"Boston", "area":89}'`)
	http.ListenAndServe(":8000", nil)
}
