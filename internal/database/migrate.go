package database

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"

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
		if strings.Contains(err.Error(), "no such file") || strings.Contains(err.Error(), "cannot find the path") {
			initFileContent = []byte("CREATE TABLE IF NOT EXISTS study_cycles (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL UNIQUE, completed_times INTEGER NOT NULL, selected BOOLEAN NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP); CREATE TABLE IF NOT EXISTS study_cycle_subjects (id INTEGER PRIMARY KEY AUTOINCREMENT, study_cycle_id INTEGER NOT NULL, name TEXT NOT NULL, max_study_hours INTEGER NOT NULL, user_studied_hours INTEGER NOT NULL, completed_times INTEGER NOT NULL, added_at TIMESTAMP NOT NULL, updated_at TIMESTAMP, FOREIGN KEY (study_cycle_id) REFERENCES study_cycles(id), UNIQUE (study_cycle_id, name));")
		} else {
			return fmt.Errorf("error on trying to read database init file: %s", err.Error())
		}
	}

	query := string(initFileContent)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("error on trying to execute init file: %s", err.Error())
	}

	return nil
}
