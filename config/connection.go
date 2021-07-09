package config

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "user=erwinwahyuramadhan dbname=jamtangan_db password= host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully conected to database!")

	return db, err
}
