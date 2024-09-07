package utils

import (
	"database/sql"
	"path"
	"strings"
)

func LocateDatabaseFile() (db *sql.DB, err error) {
	defaultPath, err := GetDefaultPath()
	if err != nil {
		return
	}

	db, err = sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		fileNotFound := strings.Contains(err.Error(), "no such file or directory")

		if !fileNotFound {
			return
		}
	}

	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return
	}

	return
}
