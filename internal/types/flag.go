package types

import (
	"flag"
	"log"
	"strings"
)

type Flag struct {
	Name        string
	Type        string
	Value       any
	Abreviation string
	Explanation string
}

func NewFlag(name, abreviation, flagType, explanation string) Flag {
	return Flag{
		Name:        name,
		Type:        flagType,
		Explanation: explanation,
		Abreviation: abreviation,
	}
}

func (f *Flag) SetFlag(flagSet *flag.FlagSet) {
	switch strings.ToUpper(f.Type) {
	case "STRING":
		f.Value = flagSet.String(f.Name, "", f.Explanation)
	case "BOOLEAN":
		f.Value = flagSet.Bool(f.Name, false, f.Explanation)
	case "INTEGER":
		f.Value = flagSet.Int(f.Name, 0, f.Explanation)
	default:
		log.Fatal("Tipo de flag inválido.")
	}
}
