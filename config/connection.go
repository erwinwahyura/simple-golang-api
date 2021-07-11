package config

import (
	"database/sql"
	"fmt"
)

var (
	dbUser    = "erwinwahyuramadhan"
	dbPass    = ""
	dbHost    = "localhost"
	dbName    = "jamtangan_db"
	pgConnStr = fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPass, dbHost)
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", pgConnStr)

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully conected to database!")

	return db, err
}
