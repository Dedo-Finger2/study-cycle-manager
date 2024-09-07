package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func SelectStudyCycle() {
	selectStudyCycle := flag.NewFlagSet("select", flag.ExitOnError)
	id := selectStudyCycle.Int("id", 0, "")

	if len(os.Args) < 3 {
		log.Fatal("this command required a flag to work.")
	}

	err := selectStudyCycle.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *id == 0 {
		log.Fatal("missing required flag id.")
	}

	err = repositories.ToggleStudyCycleSelected(*id)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	log.Printf("selected study cycle with id '%d'", *id)
}
