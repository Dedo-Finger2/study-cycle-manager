package database

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func MigrateSqlite() {
	db, err := sql.Open("sqlite3", path.Join("./", "internal", "store", "sqlite", "database.db"))
	if err != nil {
		log.Fatalf("error on trying to open sqlite connection: %s", err.Error())
		os.Exit(1)
	}

	Migrate(db)
}
