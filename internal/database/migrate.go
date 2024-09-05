package database

import (
	"database/sql"
	"log"
	"os"
	"path"
)

func Migrate(db *sql.DB) {
	defer db.Close()

	initFileContent, err := os.ReadFile(path.Join("./", "internal", "store", "sqlite", "migrations", "init.sql"))
	if err != nil {
		log.Fatalf("error on trying to read database init file: %s", err.Error())
		os.Exit(1)
	}

	query := string(initFileContent)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("error on trying to execute init file: %s", err)
		os.Exit(1)
	}
}
