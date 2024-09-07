package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func ListAllStudyCycles() {
	listAllStudyCyclesCMD := flag.NewFlagSet("list", flag.ExitOnError)
	selected := listAllStudyCyclesCMD.Bool("selected", false, "")

	err := listAllStudyCyclesCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	if *selected {
	} else {
		studyCycles, err := repositories.GetStudyCycles()
		if err != nil {
			if strings.Contains(err.Error(), "no such table") {
				log.Fatal("migrate the database first.")
			}

			log.Fatal(err)
		}

		defer fmt.Println(strings.Repeat("-", 150))

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
		defer w.Flush()

		fmt.Println(strings.Repeat("-", 150))

		fmt.Fprintln(w, "ID\t TITLE\t COMPLETED_TIMES\t SELECTED\t CREATED_AT\t UPDATED_AT")

		for _, studyCycle := range studyCycles {
			fmt.Fprintln(w, studyCycle.ID, "\t", studyCycle.Title, "\t", studyCycle.CompletedTimes, "\t", studyCycle.Selected, "\t", studyCycle.CreatedAt, "\t", studyCycle.UpdatedAt)
		}

	}
}
