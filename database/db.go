package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"danielweaver.dev/go-todo/utils"
)

func Open() *sql.DB {
	SQL, err := sql.Open("sqlite3", "./database/database.sqlite")
	utils.CheckError(err, "Error attempting to connect to database.")

	return SQL
}
