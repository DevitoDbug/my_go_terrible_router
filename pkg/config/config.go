package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func GetDBInstance() *sql.DB {
	return Db
}
func ConnectToDb() error {
	connectionString := "root:j1751502021@tcp(localhost:3306)/notesAPI_II?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	Db = db
	return nil
}
func CloseConnectionToDb() {
	if Db != nil {
		err := Db.Close()
		if err != nil {
			log.Fatalf("Error in closing the db connection: %v", err)
			return
		}
	} else {
		log.Printf("Db connection was nil")
	}
}
