package db

import (
	"database/sql"
)

func DB() (db *sql.DB) {
	db, _ = sql.Open("sqlite3", "/usr/share/sing-boxa/sing-boxa.db")
	return db
}
