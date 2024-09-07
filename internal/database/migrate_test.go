package database_test

import (
	"database/sql"
	"log"
	"os"
	"path"
	"testing"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/database"
	"github.com/Dedo-Finger2/study-cycle-manager/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

func TestSqliteDatabaseMigrate(t *testing.T) {
	defaultPath, err := utils.GetDefaultPath()
	if err != nil {
		t.Error(err)
	}

	db, err := sql.Open("sqlite3", path.Join(defaultPath, "internal", "store", "sqlite", "test_database.db"))
	if err != nil {
		log.Fatalf("error on trying to open sqlite connection: %s", err.Error())
		os.Exit(1)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Error(err)
	}
}
