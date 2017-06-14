package service

import (
	"github.com/gorilla/mux"
)

//NewRouter Main routing function
func NewRouter() *mux.Router {
	//Create an instance of gorilla router
	router := mux.NewRouter().StrictSlash(true)

	//Iterate over routes we declared in routes.go and attach them to this router
	for _, route := range routes {
		//Attach each route using a builder-like pattern to set each route up
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
