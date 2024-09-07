package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func RemoveSubject() {
	removeSubjectCMD := flag.NewFlagSet("remove", flag.ExitOnError)
	id := removeSubjectCMD.Int("id", 0, "")

	if len(os.Args) < 3 {
		log.Fatal("this command only works if you use a flag with it.")
	}

	err := removeSubjectCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error trying to parse flags: %s", err)
	}

	if *id == 0 {
		log.Fatal("expected flag --id.")
	}

	var removeConfirmation string

	fmt.Println("Do you really want to remove the subject?")
	fmt.Println("[y]es  [n]o")

	_, err = fmt.Scan(&removeConfirmation)
	if err != nil {
		log.Fatal(err)
	}

	removeConfirmation = strings.TrimSpace(removeConfirmation)
	removeConfirmation = strings.ToUpper(string(removeConfirmation[0]))

	if removeConfirmation != "Y" {
		os.Exit(1)
	}

	err = repositories.DeleteCycleSubject(*id)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	log.Println("subject removed!")
}
