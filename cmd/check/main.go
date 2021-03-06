package main

import (
	"github.com/concourse/registry-image-resource/commands"
	"github.com/fatih/color"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	color.NoColor = false

	command := commands.NewCheck(
		os.Stdin,
		os.Stderr,
		os.Stdout,
		os.Args,
	)

	err := command.Execute()
	if err != nil {
		logrus.Errorf("%s", err)
		os.Exit(1)
	}
}
