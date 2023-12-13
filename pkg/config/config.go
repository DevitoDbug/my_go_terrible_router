package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Db *sql.DB

func GetDBInstance() *sql.DB {
	return Db
}

func ConnectToDb() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set in the environment")
	}
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}
	Db = db
	return nil
}

func CloseConnectionToDb() {
	if Db != nil {
		err := Db.Close()
		if err != nil {
			log.Fatalf("Error in closing the db connection: %v", err)
		}
	} else {
		log.Println("Db connection was nil")
	}
}
