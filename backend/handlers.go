package main

import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ShowPersons(w http.ResponseWriter, r *http.Request) {
	person := Persons{
		Person{Name: "Daniel Kanczuk teste"},
		Person{Name: "Marilia Melo Teste"},
	}

	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}
}

func ShowOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["personId"]
	fmt.Fprintln(w, "Person show:", personId)
}