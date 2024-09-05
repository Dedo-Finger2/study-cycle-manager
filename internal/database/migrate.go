package database

import (
	"database/sql"
	"log"
	"os"
	"path"
)

func Migrate(db *sql.DB) {
	defer db.Close()

	initFile, err := os.ReadFile(path.Join("./../store/sqlite/migrations/init.sql"))
	if err != nil {
		log.Fatalf("error on trying to read database init file: %s", err.Error())
	}

	query := string(initFile)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("error on trying to execute init file: %s", err.Error())
	}
}
