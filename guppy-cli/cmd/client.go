package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func put() cli.Command {
	command := cli.Command{}
	command.Name = "put"
	command.Usage = "put [command]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "File Template Path",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		fmt.Println("OK")
		return nil
	}

	return command
}
