package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(25) NOT NULL,
		description VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		userId VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL
	);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
