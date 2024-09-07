package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/store/sqlite/repositories"
)

func ViewStudyCycleProgress() {
	viewStudyCycleProgressCMD := flag.NewFlagSet("view", flag.ExitOnError)

	if len(os.Args) >= 3 {
		log.Fatal("this command does not support any flags.")
	}

	err := viewStudyCycleProgressCMD.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("error on trying to parse flags: %s", err)
	}

	subjects, err := repositories.GetCurrentStudyCycleSubjects()
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

	fmt.Fprintln(w, "ID\t NAME\t COMPLETED_TIMES\t ADDED_AT\t PROGRESS")

	sort.Slice(subjects, func(i, j int) bool {
		return subjects[i].ID < subjects[j].ID
	})

	for _, subject := range subjects {
		progress := strings.Repeat("⬛", subject.UserStudiedHours) + strings.Repeat("⬜", subject.MaxStudyHours-subject.UserStudiedHours)

		fmt.Fprintln(w, strconv.Itoa(subject.ID), "\t", subject.Name, "\t", subject.CompletedTimes, "\t", subject.AddedAt, "\t", progress)
	}
}
