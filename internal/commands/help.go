package commands

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type flagSpec struct {
	Flag        string
	Type        string
	Explanation string
}

type commandSpec struct {
	Explanation string
	Flags       []flagSpec
}

func Help() {
	defer fmt.Println(strings.Repeat("-", 130))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 9, ' ', 0)
	defer w.Flush()

	commands := map[string]commandSpec{
		"add": {
			Explanation: "Adds a new subject to the current selected study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--name",
					Type:        "string",
					Explanation: "the name of the subject.",
				},
				{
					Flag:        "--max-study-hours",
					Type:        "integer",
					Explanation: "the maximum time you spend studying this subject before locking it.",
				},
			},
		},
		"create": {
			Explanation: "Creates a new study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--title",
					Type:        "string",
					Explanation: "the title of your study cycle.",
				},
			},
		},
		"delete": {
			Explanation: "Deletes a study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the study cycle to be deleted.",
				},
			},
		},
		"list": {
			Explanation: "Returns a list of all study cycles current created.",
		},
		"remove": {
			Explanation: "Removes a subject from the current selected study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the subject to be removed.",
				},
			},
		},
		"reset": {
			Explanation: "Resets the current study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the study cycle to be reseted.",
				},
			},
		},
		"select": {
			Explanation: "Selects a different study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the study cycle to be selected.",
				},
			},
		},
		"study": {
			Explanation: "Adds 1 hour to a subject in the current selected study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the subject to be studied.",
				},
			},
		},
		"un-study": {
			Explanation: "Removes 1 hour to a subject in the current selected study cycle.",
			Flags: []flagSpec{
				{
					Flag:        "--id",
					Type:        "integer",
					Explanation: "the id of the subject to be un-studied.",
				},
			},
		},
		"view": {
			Explanation: "Returns a view of all subjects in the current selected study cycle with their progress.",
		},
	}

	fmt.Fprintln(w, "COMMAND\t EXPLANATION")
	fmt.Println(strings.Repeat("-", 130))

	for command, spec := range commands {
		fmt.Fprintln(w, command, "\t", spec.Explanation)
		for _, flag := range spec.Flags {
			fmt.Fprintln(w, " ", flag.Flag, flag.Type, "\t", flag.Explanation)
		}
	}
}
