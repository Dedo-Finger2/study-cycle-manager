package database

import (
	"database/sql"
	"log"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func MigrateSqlite() {
	db, err := sql.Open("sqlite", path.Join("./../store/sqlite/database.db"))
	if err != nil {
		log.Fatalf("error on trying to open sqlite connection: %s", err.Error())
	}

	Migrate(db)
}
