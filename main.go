package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

const basePath = "/api"

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	fmt.Println(userName, password, dbName, dbHost)

	SetupRoutes(basePath)
	fmt.Println("\nWebserver is up and running on port 5000, waiting for connections...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
