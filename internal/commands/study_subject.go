package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func StudySubject() {
	studySubjectCMD := flag.NewFlagSet("study", flag.ExitOnError)
	id := studySubjectCMD.Int("id", 0, "")
	flags := []string{"id"}

	if len(os.Args) < 3 {
		log.Fatal("this command requires a flag to work. Try these: ", strings.Join(flags, ", "))
	}

	err := studySubjectCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *id == 0 {
		log.Fatal("missing required flag id.")
	}

	err = repositories.AddOneHourToSubject(*id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("studies 1 hour on the subject with id '%d'", *id)
}
