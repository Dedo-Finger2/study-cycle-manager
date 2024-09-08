package types

import (
	"flag"
	"log"
	"os"
)

type Command struct {
	Name        string
	Explanation string
	Handler     func(flagSet *flag.FlagSet)
	Flags       []Flag
}

func NewCommand(name, explanation string, handler func(flagSet *flag.FlagSet)) Command {
	return Command{
		Name:        name,
		Explanation: explanation,
		Handler:     handler,
	}
}

func (c *Command) AddFlag(name, abreviation, flagType, explanation string) {
	c.Flags = append(c.Flags, NewFlag(name, abreviation, flagType, explanation))
}

func (c Command) parseFlags(flagSet *flag.FlagSet) {
	if err := flagSet.Parse(os.Args[2:]); err != nil {
		log.Fatalf("Erro ao tentar fazer parse das flags para o comando '%s': %v", c.Name, err)
	}
}

func (c *Command) Run() {
	command := flag.NewFlagSet(c.Name, flag.ExitOnError)
	for _, f := range c.Flags {
		f.SetFlag(command)
	}
	c.parseFlags(command)
	c.Handler(command)
}
