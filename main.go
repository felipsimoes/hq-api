package main

import (
	"fmt"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	SetupRoutes(basePath)
	fmt.Println("Webserver is up and running, waiting for connections...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
