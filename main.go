package main

import (
	"fmt"

	"./services"
)

var appName = "gorun"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8080")
}
