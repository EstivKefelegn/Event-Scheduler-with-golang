package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB () {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to databse")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}


func CreateTable () {

	usersTable := `
		CREATE TABLE IF NPT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(usersTable)
	
	if err != nil {
		panic("Couldn't create the table")
	}

	eventsEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(eventsEventTable)

	if err != nil {
		panic("Cant create the table" + err.Error())
	}
}