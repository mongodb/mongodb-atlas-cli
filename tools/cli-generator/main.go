package main

import (
	"flag"
	"log"
)

func main() {
	overwrite := flag.Bool("overwrite", false, "regenerate files")
	flag.Parse()

	cli, err := newCli(*overwrite)
	if err != nil {
		log.Fatal(err)
	}
	err = cli.generateCli()
	if err != nil {
		log.Fatal(err)
	}
}
