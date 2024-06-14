package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

type Save struct {
	ID      int
	Content string
}

func initDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func createTables(db *sql.DB) error {
	createSaveTable := `
    CREATE TABLE IF NOT EXISTS Save (
	    ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Content TEXT
    );`

	_, err := db.Exec(createSaveTable)
	if err != nil {
		log.Printf("Error creating User table: %v", err)
		return err
	}
	log.Println("User table created successfully or already exists")

	return nil

}
