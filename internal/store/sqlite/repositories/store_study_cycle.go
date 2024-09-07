package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func StoreStudyCycle(title string) error {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return err
	}

	defer db.Close()

	createdAt := time.Now().Local()
	selected := false
	completedTimes := 0
	updatedAt := time.Now().Local()

	selectStmt, err := db.Prepare(`
		SELECT id FROM study_cycles
		WHERE title = ?
	`)
	if err != nil {
		return err
	}

	defer selectStmt.Close()

	var nameAlreadyInUse int

	err = selectStmt.QueryRow(title).Scan(&nameAlreadyInUse)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if nameAlreadyInUse != 0 {
		return fmt.Errorf("title '%s' is already in use.", title)
	}

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
