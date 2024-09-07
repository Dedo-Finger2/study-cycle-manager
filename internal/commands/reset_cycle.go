package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func ResetCycle() {
	resetCycleCMD := flag.NewFlagSet("reset", flag.ExitOnError)
	partial := resetCycleCMD.Bool("partial", false, "")

	err := resetCycleCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	var resetConfirmation string

	fmt.Println("Do you really want to reset the current study cycle?")
	fmt.Println("[y]es  [n]o")

	_, err = fmt.Scan(&resetConfirmation)
	if err != nil {
		log.Fatal(err)
	}

	resetConfirmation = strings.TrimSpace(resetConfirmation)
	resetConfirmation = strings.ToUpper(string(resetConfirmation[0]))

	if resetConfirmation != "Y" {
		os.Exit(1)
	}

	completed, err := repositories.CheckCycleCompleted()
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	if completed && !*partial {
		err = repositories.AddOneToCycleCompletedTimes()
		if err != nil {
			if strings.Contains(err.Error(), "no such table") {
				log.Fatal("migrate the database first.")
			}

			log.Fatalf("error trying to add one hour to completed times in cycle: %s", err)
		}

		err = repositories.AddOneToCycleSubjectsCompletedTimes()
		if err != nil {
			if strings.Contains(err.Error(), "no such table") {
				log.Fatal("migrate the database first.")
			}

			log.Fatalf("error trying to add one hour to completed times in cycle's subjects: %s", err)
		}

		err = repositories.ResetCyclesubjectsUserStudiedHours()
		if err != nil {
			if strings.Contains(err.Error(), "no such table") {
				log.Fatal("migrate the database first.")
			}

			log.Fatalf("error trying to reset cycle's subjects user studied hours: %s", err)
		}
	} else {
		log.Println("cycle is not completed yet.")
	}
}
