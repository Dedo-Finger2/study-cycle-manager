package repositories

import (
	"database/sql"
	"path"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func ResetCyclesubjectsUserStudiedHours() error {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return err
	}

	defer db.Close()

	studyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
    UPDATE study_cycle_subjects
    SET user_studied_hours = 0
    WHERE study_cycle_id = ?
  `)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studyCycleID)
	if err != nil {
		return err
	}

	return nil
}
