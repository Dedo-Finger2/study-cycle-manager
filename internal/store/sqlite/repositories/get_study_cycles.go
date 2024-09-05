package repositories

import (
	"database/sql"
	"path"
	"time"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

type StudyCycle struct {
	ID             int
	Title          string
	CompletedTimes int
	Selected       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func GetStudyCycles() (studyCycles []StudyCycle, err error) {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM study_cycles")
	if err != nil {
		return
	}

	for rows.Next() {
		var studyCycle StudyCycle = StudyCycle{}

		err = rows.Scan(&studyCycle.ID, &studyCycle.Title, &studyCycle.CompletedTimes, &studyCycle.Selected, &studyCycle.CreatedAt, &studyCycle.UpdatedAt)
		if err != nil {
			return
		}

		studyCycles = append(studyCycles, studyCycle)
	}

	return
}
