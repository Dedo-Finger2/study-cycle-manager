package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/types"
)

type Cli struct {
	Commands []*types.Command
	Flags    []types.Flag
}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) AddCommand(command *types.Command) {
	c.Commands = append(c.Commands, command)
}

func (c *Cli) Run() {
	for _, command := range c.Commands {
		c.Flags = append(c.Flags, command.Flags...)
	}
	userInput := os.Args[1]
	commandFound := false
	for _, command := range c.Commands {
		if strings.EqualFold(userInput, command.Name) {
			command.Run()
			commandFound = true
		}
	}
	if !commandFound {
		fmt.Printf("Invalid command '%s'. Try using --help to see avaliable commands.", userInput)
		os.Exit(1)
	}
}
