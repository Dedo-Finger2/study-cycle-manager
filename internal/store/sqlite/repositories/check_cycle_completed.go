package repositories

import (
	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func CheckCycleCompleted() (completed bool, err error) {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return
	}

	defer db.Close()

	studyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return
	}

	selectStmt, err := db.Prepare(`
    SELECT
    SUM(max_study_hours) as max_study_hours_sum,
    SUM(user_studied_hours) as user_studied_hours_sum
    FROM study_cycle_subjects
    WHERE study_cycle_id = ?
  `)

	type StudyHoursSum struct {
		max_study_hours    int
		user_studied_hours int
	}

	var studyHoursSum StudyHoursSum = StudyHoursSum{}

	row := selectStmt.QueryRow(studyCycleID)
	err = row.Scan(&studyHoursSum.max_study_hours, &studyHoursSum.user_studied_hours)
	if err != nil {
		return
	}

	if studyHoursSum.max_study_hours == studyHoursSum.user_studied_hours {
		completed = true
		return
	}

	return
}
