package cmd

import (
	"github.com/sofyan48/guppy/guppy-cli/libs"
	"github.com/sofyan48/guppy/guppy-cli/utils"
	"github.com/urfave/cli"
)

// CLIMapping ...
type CLIMapping struct {
	Utils utils.UtilsInterface
	Lib   libs.LibraryInterface
}

// CLIMappingHandler ...
func CLIMappingHandler() *CLIMapping {
	return &CLIMapping{
		Utils: utils.UtilsHandler(),
		Lib:   libs.LibraryHandler(),
	}
}

// ArgsMapping object mapping
type ArgsMapping struct {
	EnvPath      string
	TemplatePath string
	Key          string
	Value        string
	Encryption   bool
}

// Args Glabal Acces args command
var Args ArgsMapping
var app *cli.App

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	handler := CLIMappingHandler()
	app.Commands = []cli.Command{
		handler.put(),
		handler.get(),
	}
	return app
}

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
