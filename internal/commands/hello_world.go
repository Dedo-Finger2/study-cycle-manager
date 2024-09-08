package commands

import (
	"flag"
	"fmt"

	"github.com/Dedo-Finger2/study-cycle-manager/internal/types"
)

var HelloWorldCMD = types.NewCommand("hello-world", "something", handler)

func init() {
	HelloWorldCMD.AddFlag("teste", "t", "string", "explicação")
}

func handler(fs *flag.FlagSet) {
	fmt.Println(fs.Lookup("teste").Value.String())
	fmt.Println("Hello, World!")
}
