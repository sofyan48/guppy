package main

import (
	"os"

	"github.com/sofyan48/guppy/guppy-cli/cmd"
)

func main() {
	app := cmd.AppCommands()
	app.Run(os.Args)
}
