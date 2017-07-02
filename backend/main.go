package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ShowPersons(w http.ResponseWriter, r *http.Request) {
	session, _ := mgo.Dial("localhost:27017")
	c := session.DB("simpleCRUD").C("person")

	person := Persons{}

	err := c.Find(bson.M{}).All(&person)

	if err != nil {
		log.Println(err)
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

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

type Person struct {
	Name      string    `json:"name" bson:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Persons []Person

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/////again!
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ShowPersons",
		"GET",
		"/person",
		ShowPersons,
	},
	Route{
		"ShowOnePerson",
		"GET",
		"/person/{personId}",
		ShowOnePerson,
	},
}
