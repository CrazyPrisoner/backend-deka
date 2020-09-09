package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func ConnectPsql() error {

	err := godotenv.Load() //Load .env file
	if err != nil {
		return err
	}

	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbName := os.Getenv("dbName")
	dbHost := os.Getenv("dbHost")
	dbPort := os.Getenv("dbPort")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass) //Build connection string

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		return err
	}
	db = conn
	// db.LogMode(true)
	fmt.Println("PostgreSQL connection established!")

	return nil
}

func Close() error {
	return db.Close()
}

func GetDB() *gorm.DB {
	return db
}
