package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"hq-collections.com/collection"
	"hq-collections.com/database"
	"hq-collections.com/volume"
)

const basePath = "/api"

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	db := database.Setup()
	migrateAndSeed(db)

	SetupRoutes(basePath)

	fmt.Println("\nWebserver is up and running on port 5000, waiting for connections...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func seed(db *gorm.DB) {
	seedCollections(*db)
	seedVolumes(*db)
}

func seedCollections(db gorm.DB) {
	db.Exec("DELETE FROM collections")

	fmt.Print("Seeding collections... ")
	fileName := "data/collections.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	file, _ := ioutil.ReadFile(fileName)
	collectionList := make([]collection.Collection, 0)
	err = json.Unmarshal([]byte(file), &collectionList)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(collectionList); i++ {
		db.Create(&collectionList[i])
	}

	fmt.Println("done.")
}

func seedVolumes(db gorm.DB) {
	db.Exec("DELETE FROM volumes")

	fmt.Print("Seeding volumes... ")
	fileName := "data/volumes.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	file, _ := ioutil.ReadFile(fileName)
	volumesList := make([]volume.Volume, 0)
	err = json.Unmarshal([]byte(file), &volumesList)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(volumesList); i++ {
		db.Create(&volumesList[i])
	}

	fmt.Println("done.")
}

func migrateAndSeed(db *gorm.DB) {
	fmt.Println("Migrating...")
	db.AutoMigrate(&collection.Collection{}, &volume.Volume{})
	seed(db)
}
