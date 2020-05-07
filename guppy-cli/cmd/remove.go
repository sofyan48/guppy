package cmd

import (
	"fmt"

	"github.com/sofyan48/guppy/guppy/entity"
	"github.com/urfave/cli"
)

func (handler *CLIMapping) remove() cli.Command {
	command := cli.Command{}
	command.Name = "rm"
	command.Usage = "rm [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Path Key",
			Destination: &Args.Key,
			Required:    true,
		},
	}
	command.Action = func(c *cli.Context) error {
		client, err := handler.Lib.GetClients(Args.EnvPath)
		if err != nil {
			return err
		}
		parameters := &entity.Parameters{}
		parameters.Path = Args.Key
		result, err := client.Del(parameters)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	}

	return command
}
