package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/database"
)

func MigrateCMD() {
	migrateCMD := flag.NewFlagSet("migrate", flag.ExitOnError)
	sqlite := migrateCMD.Bool("sqlite", false, "migrates with sqlite db.")
	flags := []string{"sqlite"}
	flagFound := false

	if len(os.Args) < 3 {
		log.Fatal("missing required flag. Choose one of these: ", strings.Join(flags, ", "))
	}

	err := migrateCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err.Error())
	}

	for _, flag := range os.Args[2:] {
		for _, validFlag := range flags {
			if strings.Trim(strings.ReplaceAll(flag, "-", " "), " ") == validFlag {
				flagFound = true
			}
		}
	}

	if !flagFound {
		log.Fatal("expecting any flag of these: ", strings.Join(flags, ", "))
	}

	switch {
	case *sqlite:
		err := database.MigrateSqlite()
		if err != nil {
			log.Fatalf("error trying to migrate database: %s", err)
		}
		fmt.Println("database migrated!")
	default:
		log.Fatal("invalid database name.")
	}
}
