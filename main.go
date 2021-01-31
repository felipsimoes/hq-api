package main

import (
	"fmt"
	"log"
	"net/http"

	"hq-collections.com/collection"
	"hq-collections.com/volume"
)

const basePath = "/api"

func main() {
	collection.SetupRoutes(basePath)
	volume.SetupRoutes(basePath)
	fmt.Println("Webserver is up and running, waiting for connections...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
