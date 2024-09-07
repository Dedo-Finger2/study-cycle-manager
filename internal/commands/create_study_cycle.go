package commands

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func CreateStudyCycle() {
	createStudyCycleCMD := flag.NewFlagSet("create", flag.ExitOnError)
	title := createStudyCycleCMD.String("title", "", "")
	flags := []string{"title"}

	if len(os.Args) < 3 {
		log.Fatal("missing required flag. Choose one of these: ", strings.Join(flags, ", "))
	}

	err := createStudyCycleCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *title == "" {
		log.Fatal("expecting a tittle, got none.")
	}

	formattedTitle := strings.ToUpper(string(string(*title)[0])) + strings.ToLower(string(string(*title)[1:]))

	err = repositories.StoreStudyCycle(formattedTitle)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			log.Fatal("migrate the database first.")
		}

		log.Fatalf("error trying to save study cycle: %s", err)
	}

	log.Printf("new study cycle named '%s' created!", formattedTitle)
}
