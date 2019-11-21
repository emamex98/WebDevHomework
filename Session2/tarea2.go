package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"log"
)

type Event struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Date  string `json:"date:`
}

var Events []Event

func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Events {
		if strconv.Itoa(item.Id) + "" == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func main() {
	Events = append(Events, Event{Id: 100, Name: "Changemaker Day", Date: "October 10, 2019"})
	Events = append(Events, Event{Id: 101, Name: "CodeGDL", Date: "October 12, 2019"})

	router := mux.NewRouter()

	router.HandleFunc("/events", GetEvents).Methods("GET")
	router.HandleFunc("/events/{id}", GetEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
