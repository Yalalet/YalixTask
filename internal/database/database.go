package database

import (
	"database/sql"
	"log"
)

func Connect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Проверьте подключение
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
