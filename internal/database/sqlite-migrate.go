package database

import (
	"log"
	"os"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

func MigrateSqlite() error {
	db, err := utils.LocateDatabaseFile()
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
