package db

import (
	"database/sql"
)

func DB() (db *sql.DB) {
	db, _ = sql.Open("sqlite3", "~/sing-boxa.db")
	return db
}
