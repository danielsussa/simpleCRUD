package main

import (
	"net/http"
)

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