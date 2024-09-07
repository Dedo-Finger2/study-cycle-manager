package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func DeleteCycle(id int) error {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return err
	}

	defer db.Close()

	selectStmt, err := db.Prepare(`
	  SELECT id
	  FROM study_cycles
	  WHERE id = ? 
	`)
	if err != nil {
		return err
	}

	var studyCycleExists int

	err = selectStmt.QueryRow(id).Scan(&studyCycleExists)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("study cycle with id '%d' was not found.", id)
	}
	if err != nil {
		return err
	}

	cycleSubjectsDeleteStmt, err := db.Prepare(`
    DELETE
    FROM study_cycle_subjects
    WHERE study_cycle_id = ?
  `)
	if err != nil {
		return err
	}

	defer cycleSubjectsDeleteStmt.Close()

	_, err = cycleSubjectsDeleteStmt.Exec(id)
	if err != nil {
		return err
	}

	cycleDeleteStmt, err := db.Prepare(`
    DELETE
    FROM study_cycles
    WHERE id = ?
  `)
	if err != nil {
		return err
	}

	defer cycleDeleteStmt.Close()

	_, err = cycleDeleteStmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
