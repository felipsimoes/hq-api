package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const basePath = "/api"

func connectionString() string {
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	fmt.Println(userName, password, dbName, dbHost)

	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, userName, dbName, password)
}

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	db, err := gorm.Open("postgres", connectionString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to the database succeeded")

	SetupRoutes(basePath)
	fmt.Println("\nWebserver is up and running on port 5000, waiting for connections...")
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
