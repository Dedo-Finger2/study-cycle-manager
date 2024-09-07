package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func AddSubject() {
	addSubjectCMD := flag.NewFlagSet("add", flag.ExitOnError)
	name := addSubjectCMD.String("name", "", "")
	maxStudyHours := addSubjectCMD.Int("max-study-hours", 0, "")

	err := addSubjectCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *name == "" {
		log.Fatal("missing required flag '-name'.")
	}

	if *maxStudyHours == 0 {
		log.Fatal("missing required flag '-max-study-hours'.")
	}

	formattedTitle := strings.ReplaceAll(strings.ToLower(strings.Trim(*name, " ")), " ", "-")

	studyCycleID, err := repositories.GetSelectedStudyCycleID()
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	err = repositories.StoreSubject(formattedTitle, studyCycleID, *maxStudyHours)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatal(err)
	}

	log.Printf("subject '%s' added!", formattedTitle)
}
