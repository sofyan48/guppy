package cmd

import (
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	EnvPath      string
	TemplatePath string
	ID           string
}

// Args Glabal Acces args command
var Args ArgsMapping

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "guppy"
	app.Usage = "guppy [command]"
	app.Author = "sofyan48"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.1.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Load environtment config from `FILE`",
			Destination: &Args.EnvPath,
		},
	}
	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		// service is here
	}
	return app
}
