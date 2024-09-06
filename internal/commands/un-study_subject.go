package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func UnStudySubject() {
	unstudySubjectCMD := flag.NewFlagSet("un-study", flag.ExitOnError)
	id := unstudySubjectCMD.Int("id", 0, "")
	flags := []string{"id"}

	if len(os.Args) < 3 {
		log.Fatal("this command requires a flag to work. Try these: ", strings.Join(flags, ", "))
	}

	err := unstudySubjectCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *id == 0 {
		log.Fatal("missing required flag id.")
	}

	err = repositories.RemoveOneHourFromSubject(*id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("removed 1 hour on the subject with id '%d'", *id)
}
