package database

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func Migrate(db *sql.DB) error {
	defer db.Close()

	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	initFileContent, err := os.ReadFile(path.Join(defaultPath, "internal", "store", "sqlite", "migrations", "init.sql"))
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
