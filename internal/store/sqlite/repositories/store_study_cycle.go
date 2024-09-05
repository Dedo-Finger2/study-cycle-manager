package repositories

import (
	"database/sql"
	"path"
	"time"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func StoreStudyCycle(title string) error {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return err
	}

	defer db.Close()

	createdAt := time.Now().Local()
	selected := false
	completedTimes := 0
	updatedAt := time.Now().Local()

	stmt, err := db.Prepare("INSERT INTO study_cycles (title, completed_times, selected, created_at, updated_at) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title, completedTimes, selected, createdAt, updatedAt)
	if err != nil {
		return err
	}

	return nil
}
