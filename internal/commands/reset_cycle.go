package commands

import (
	"flag"
	"log"
	"os"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func ResetCycle() {
	resetCycleCMD := flag.NewFlagSet("reset", flag.ExitOnError)
	partial := resetCycleCMD.Bool("partial", false, "")

	err := resetCycleCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	completed, err := repositories.CheckCycleCompleted()
	if err != nil {
		log.Fatal(err)
	}

	if completed && !*partial {
		err = repositories.AddOneToCycleCompletedTimes()
		if err != nil {
			log.Fatalf("error trying to add one hour to completed times in cycle: %s", err)
		}

		err = repositories.AddOneToCycleSubjectsCompletedTimes()
		if err != nil {
			log.Fatalf("error trying to add one hour to completed times in cycle's subjects: %s", err)
		}

		err = repositories.ResetCyclesubjectsUserStudiedHours()
		if err != nil {
			log.Fatalf("error trying to reset cycle's subjects user studied hours: %s", err)
		}
	} else {
		log.Println("cycle is not completed yet.")
	}
}
