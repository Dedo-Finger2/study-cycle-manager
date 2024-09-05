package main

import (
	"log"
	"os"
	"strings"

	cmd "github.com/Dedo-Finger2/study-cycle-manager/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected a command, found none.")
	}

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
	case "SELECT":
		cmd.SelectStudyCycle()
	default:
		log.Printf("command '%s' not found.", command)
		os.Exit(1)
	}
}
