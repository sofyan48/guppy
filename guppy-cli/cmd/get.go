package cmd

import (
	"log"

	"github.com/urfave/cli"
)

func (handler *CLIMapping) get() cli.Command {
	command := cli.Command{}
	command.Name = "get"
	command.Usage = "get [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Path Key",
			Destination: &Args.Key,
		},
	}
	command.Action = func(c *cli.Context) error {
		client, err := handler.Lib.GetClients(Args.EnvPath)
		if err != nil {
			return err
		}
		result, err := client.Get(Args.Key)
		if err != nil {
			return err
		}
		log.Println("Path: ", string(result.Kvs[0].Key))
		log.Println("Value: ", string(result.Kvs[0].Value))
		log.Println("Create Revision: ", result.Kvs[0].CreateRevision)
		log.Println("Mod Revision: ", result.Kvs[0].ModRevision)
		return nil
	}

	return command
}