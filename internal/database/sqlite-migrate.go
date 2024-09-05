package database

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

func MigrateSqlite() error {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "database.db"))
	if err != nil {
		log.Fatalf("error on trying to open sqlite connection: %s", err.Error())
		os.Exit(1)
	}

	err = Migrate(db)
	if err != nil {
		return err
	}

	return nil
}
