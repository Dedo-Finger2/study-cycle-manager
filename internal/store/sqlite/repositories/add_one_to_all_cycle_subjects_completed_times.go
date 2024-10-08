package repositories

import (
	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func AddOneToCycleSubjectsCompletedTimes() error {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return err
	}

	defer db.Close()

	studyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return err
	}

	type SubjectCompletedTimesAndID struct {
		ID             int
		CompletedTimes int
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	selectStmt, err := tx.Prepare(`
    SELECT completed_times, id
    FROM study_cycle_subjects
    WHERE study_cycle_id = ?
  `)
	if err != nil {
		return err
	}

	rows, err := selectStmt.Query(studyCycleID)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var subjectCompletedTimesAndID SubjectCompletedTimesAndID = SubjectCompletedTimesAndID{}

		err = rows.Scan(&subjectCompletedTimesAndID.CompletedTimes, &subjectCompletedTimesAndID.ID)
		if err != nil {
			return err
		}

		stmt, err := tx.Prepare(`
      UPDATE study_cycle_subjects
      SET completed_times = ?
      WHERE id = ?
    `)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(subjectCompletedTimesAndID.CompletedTimes+1, subjectCompletedTimesAndID.ID)
		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
