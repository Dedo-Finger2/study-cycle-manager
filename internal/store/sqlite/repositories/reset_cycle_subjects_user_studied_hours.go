package repositories

import (
	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func ResetCyclesubjectsUserStudiedHours() error {
	db, err := utils.LocateDatabaseFile()
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
