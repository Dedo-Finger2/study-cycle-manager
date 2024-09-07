package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func DeleteStudyCycle() {
	deleteStudycycleCMD := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteStudycycleCMD.Int("id", 0, "")

	if len(os.Args) < 3 {
		log.Fatal("this command needs at least one flag to work.")
	}

	err := deleteStudycycleCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error trying to parse flags: %s", err)
	}

	if *id == 0 {
		log.Fatal("expected flag --id to have a value other than 0.")
	}

	err = repositories.DeleteCycle(*id)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	log.Println("study cycle deleted!")
}
