package repositories

import (
	"database/sql"
	"errors"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func GetSelectedStudyCycleID() (studyCycleID int, err error) {
	db, err := utils.LocateDatabaseFile()
	if err != nil {
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT id FROM study_cycles WHERE selected = true")
	if row.Err() != nil {
		err = row.Err()
		return
	}

	err = row.Scan(&studyCycleID)
	if errors.Is(err, sql.ErrNoRows) {
		err = errors.New("select a study cycle first.")
		return
	}
	if err != nil {
		return
	}

	return
}
