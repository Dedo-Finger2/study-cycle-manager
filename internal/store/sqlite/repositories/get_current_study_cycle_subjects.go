package repositories

import (
	"database/sql"
	"path"
	"time"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

type StudyCycleSubject struct {
	ID               int
	Name             string
	MaxStudyHours    int
	UserStudiedHours int
	CompletedTimes   int
	AddedAt          time.Time
}

func GetCurrentStudyCycleSubjects() (subjects []StudyCycleSubject, err error) {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return
	}

	defer db.Close()

	studyCycleID, err := GetSelectedStudyCycleID()
	if err != nil {
		return
	}

	stmt, err := db.Prepare(`
    SELECT
    id, name, max_study_hours, user_studied_hours, completed_times, added_at
    FROM study_cycle_subjects
    WHERE study_cycle_id = ?
  `)

	rows, err := stmt.Query(studyCycleID)
	if err != nil {
		return
	}

	for rows.Next() {
		var subject StudyCycleSubject = StudyCycleSubject{}

		err = rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.MaxStudyHours,
			&subject.UserStudiedHours,
			&subject.CompletedTimes,
			&subject.AddedAt,
		)
		if err != nil {
			return
		}

		subjects = append(subjects, subject)
	}

	return
}
