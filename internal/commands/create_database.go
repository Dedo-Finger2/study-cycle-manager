package commands

import (
	"log"
	"os"
)

func CreateDatabase() {
	_, err := os.Create("./database.db")
	if err != nil {
		log.Fatal(err)
	}
}
