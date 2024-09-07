package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func DeleteCycleSubject(id int) error {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return err
	}

	defer db.Close()

	studyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return err
	}

	selectStmt, err := db.Prepare(`
	  SELECT id
	  FROM study_cycle_subjects
	  WHERE id = ? AND study_cycle_id = ?
	`)
	if err != nil {
		return err
	}

	var subjectExists int

	err = selectStmt.QueryRow(id, studyCycleID).Scan(&subjectExists)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("subject with id '%d' was not found in the selected cycle.", id)
	}
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
    DELETE
    FROM study_cycle_subjects
    WHERE study_cycle_id = ?
    AND id = ?
  `)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(studyCycleID, id)
	if err != nil {
		return err
	}

	return nil
}
