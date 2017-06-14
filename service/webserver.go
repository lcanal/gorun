package service

import (
	"log"
	"net/http"
)

//StartWebServer Main webserver start function
func StartWebServer(port string) {
	log.Println("Starting HTTP interface at " + port)

	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil) //Goroutine blocks here

	if err != nil {
		log.Println("An error occured while attempting to listen on port " + port)
		log.Println("Error : " + err.Error())
	}
}
