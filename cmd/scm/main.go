package main

import (
	"flag"
	"log"
	"os"
	"strings"

	cmd "github.com/Dedo-Finger2/study-cycle-manager/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected a command, found none.")
	}

	flag.Usage = func() {
		cmd.Help()
	}
	flag.Parse()

	command := os.Args[1]

	switch strings.ToUpper(command) {
	case "MIGRATE":
		cmd.MigrateCMD()
	case "CREATE":
		cmd.CreateStudyCycle()
	case "LIST":
		cmd.ListAllStudyCycles()
	case "ADD":
		cmd.AddSubject()
	case "VIEW":
		cmd.ViewStudyCycleProgress()
	case "UN-STUDY":
		cmd.UnStudySubject()
	case "DELETE":
		cmd.DeleteStudyCycle()
	case "REMOVE":
		cmd.RemoveSubject()
	case "RESET":
		cmd.ResetCycle()
	case "STUDY":
		cmd.StudySubject()
	case "SELECT":
		cmd.SelectStudyCycle()
	default:
		log.Printf("command '%s' not found.", command)
		os.Exit(1)
	}
}
