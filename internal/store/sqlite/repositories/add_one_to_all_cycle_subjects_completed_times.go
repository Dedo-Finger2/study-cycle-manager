package repositories

import (
	"database/sql"
	"path"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func AddOneToCycleSubjectsCompletedTimes() error {
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

	type SubjectCompletedTimesAndID struct {
		ID             int
		CompletedTimes int
	}

	selectStmt, err := db.Prepare(`
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

	for rows.Next() {
		var subjectCompletedTimesAndID SubjectCompletedTimesAndID = SubjectCompletedTimesAndID{}

		err = rows.Scan(&subjectCompletedTimesAndID.CompletedTimes, &subjectCompletedTimesAndID.ID)
		if err != nil {
			return err
		}

		stmt, err := db.Prepare(`
      UPDATE study_cycle_subjects
      SET completed_times = ?
      WHERE id = ?
    `)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(subjectCompletedTimesAndID.CompletedTimes+1, subjectCompletedTimesAndID.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
