package main

import (
	"log"
	"os"

	"github.com/logsquaredn/static-resource/command"
)

func main() {
	command := commands.NewStaticResource(
		os.Stdin,
		os.Stderr,
		os.Stdout,
		os.Args,
	)

	err := command.Out()
	if err != nil {
		log.Fatal(err)
	}
}
