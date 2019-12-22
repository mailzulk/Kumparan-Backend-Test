package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import postgresql
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

// Connect - Connect to database with the parameter from .env
func Connect() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", dbHost, username, dbName, password, dbPort)

	// Please define your user name and password for my sql.
	d, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB - Get the database
func GetDB() *gorm.DB {
	return db
}
