package repositories

import (
	"database/sql"
	"errors"
	"path"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
)

func ToggleStudyCycleSelected(id int) error {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		return err
	}

	defer db.Close()

	selectStmt, err := db.Prepare("SELECT selected FROM study_cycles WHERE id = ?")
	if err != nil {
		return err
	}

	var studyCycleFoundSelected bool

	row := selectStmt.QueryRow(id)
	err = row.Scan(&studyCycleFoundSelected)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("study cycle not found.")
		}
		return err
	}

	if studyCycleFoundSelected {
		return errors.New("study cycle is already selected.")
	}

	_, err = db.Exec("UPDATE study_cycles SET selected = false WHERE selected = true")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE study_cycles SET selected = true WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
