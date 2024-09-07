package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func RemoveOneHourFromSubject(id int) error {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return err
	}

	defer db.Close()

	currentStudyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return err
	}
	if currentStudyCycleID == 0 {
		return errors.New("no study cycle selected.")
	}

	var subjectStudyHours struct {
		lastUserStudiedHours int
		maxStudyHours        int
	}

	selectStmt, err := db.Prepare("SELECT user_studied_hours, max_study_hours FROM study_cycle_subjects WHERE id = ? AND study_cycle_id = ?")
	if err != nil {
		return err
	}

	row := selectStmt.QueryRow(id, currentStudyCycleID)
	if row.Err() != nil {
		return row.Err()
	}

	err = row.Scan(&subjectStudyHours.lastUserStudiedHours, &subjectStudyHours.maxStudyHours)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("subject with id '%d' was not found.", id)
	}
	if err != nil {
		return err
	}

	if subjectStudyHours.lastUserStudiedHours == 0 {
		return errors.New("you cannot remove 1 study hour from this subject (subject has 0 studied hours).")
	}

	updateStmt, err := db.Prepare("UPDATE study_cycle_subjects SET user_studied_hours = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = updateStmt.Exec(subjectStudyHours.lastUserStudiedHours-1, id)
	if err != nil {
		return err
	}

	return nil
}
