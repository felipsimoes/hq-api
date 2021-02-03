package main

import (
	"fmt"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	SetupRoutes(basePath)
	fmt.Println("\nWebserver is up and running on port 5000, waiting for connections...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
