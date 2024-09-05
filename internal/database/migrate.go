package database

import (
	"database/sql"
	"fmt"
	"os"
	"path"
)

func Migrate(db *sql.DB) error {
	defer db.Close()

	initFileContent, err := os.ReadFile(path.Join("./", "internal", "store", "sqlite", "migrations", "init.sql"))
	if err != nil {
		return fmt.Errorf("error on trying to read database init file: %s", err.Error())
	}

	query := string(initFileContent)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("error on trying to execute init file: %s", err.Error())
	}

	return nil
}
