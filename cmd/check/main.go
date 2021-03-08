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

	err := command.Check()
	if err != nil {
		log.Fatal(err)
	}
}
