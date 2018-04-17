package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
var people []Person

func getPersonEndPoint(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func getPeopleEndPoint(w http.ResponseWriter, req *http.Request)  {
	json.NewEncoder(w).Encode(people)

}
func createPersonEndPoint(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func deletePersonEndPoint(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "3", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "6", Firstname: "Maria", Lastname: "Raboy"})

	router.HandleFunc("/people", getPeopleEndPoint).Methods("GET")
	router.HandleFunc("/person/{id}", getPersonEndPoint).Methods("GET")
	router.HandleFunc("/person", createPersonEndPoint).Methods("POST")
	router.HandleFunc("/person/{id}", deletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345",router))
}
