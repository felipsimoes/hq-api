package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Setup() *gorm.DB {
	logLevel := logger.Error
	if os.Getenv("debug") != "" {
		logLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err.Error())
	}
	// defer postgres.Close()

	DB = db

	// databasePing(db)

	return db
}

func connectionString() string {
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	fmt.Println(userName, password, dbName, dbHost)

	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, userName, dbName, password)
}

// func databasePing(db *gorm.DB) {
// 	err := postgres.Ping()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println("Connection to the database succeeded")
// }
