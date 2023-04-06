package main

import (
	"log"
)

func main() {
	cli, err := newCli()
	if err != nil {
		log.Fatal(err)
	}
	err = cli.generateCli()
	if err != nil {
		log.Fatal(err)
	}
}
