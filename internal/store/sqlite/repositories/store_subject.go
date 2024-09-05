package repositories

import (
	"database/sql"
	"errors"
	"path"
	"time"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func StoreSubject(name string, studyCycleID, maxStudyHours int) error {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return err
	}

	defer db.Close()

	selectStmt, err := db.Prepare("SELECT id FROM study_cycle_subjects WHERE name = ? AND study_cycle_id = ?")
	if err != nil {
		return err
	}

	var studyCycleSubjectID int

	row := selectStmt.QueryRow(name, studyCycleID)
	err = row.Scan(&studyCycleSubjectID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if studyCycleSubjectID != 0 {
		return errors.New("subject name already registered.")
	}

	userStudiedHours := 0
	completedTimes := 0
	addedAt := time.Now().Local()
	updatedAt := time.Now().Local()

	stmt, err := db.Prepare(`
    INSERT INTO study_cycle_subjects
    (study_cycle_id, name, max_study_hours, user_studied_hours, completed_times, added_at, updated_at)
    VALUES (?,?,?,?,?,?,?)
  `)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studyCycleID, name, maxStudyHours, userStudiedHours, completedTimes, addedAt, updatedAt)
	if err != nil {
		return err
	}

	return nil
}
