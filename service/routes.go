package service

import (
	"net/http"
)

//Route Main structure of a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes Full array of known routes
type Routes []Route

//Initialize our routes

var routes = Routes{

	Route{
		"GetAccount", //Name
		"GET",        //HTTP Method
		"/accounts/{accountId}", // Route Pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Context-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
