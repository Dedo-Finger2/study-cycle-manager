package repositories

import (
	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func AddOneToCycleCompletedTimes() error {
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
    SELECT completed_times
    FROM study_cycles
    WHERE id = ?
  `)
	if err != nil {
		return err
	}

	var studyCycleCompletedTimes int

	row := selectStmt.QueryRow(studyCycleID)
	err = row.Scan(&studyCycleCompletedTimes)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
    UPDATE study_cycles
    SET completed_times = ?
    WHERE id = ?
  `)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studyCycleCompletedTimes+1, studyCycleID)
	if err != nil {
		return err
	}

	return nil
}
