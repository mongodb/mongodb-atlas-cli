package main

import (
	"log"

	"github.com/mongodb/mongocli/internal/cli/root"
	"github.com/spf13/cobra/doc"
)

func main() {
	var profile string
	mongocli := root.Builder(&profile, []string{})

	err := doc.GenReSTTree(mongocli, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
