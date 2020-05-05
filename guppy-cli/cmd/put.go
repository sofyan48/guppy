package cmd

import (
	"log"

	"github.com/urfave/cli"
)

func (handler *CLIMapping) put() cli.Command {
	command := cli.Command{}
	command.Name = "put"
	command.Usage = "put [command]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "File Template Path",
			Destination: &Args.TemplatePath,
		},
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Key for path",
			Destination: &Args.Key,
		},

		cli.StringFlag{
			Name:        "value, v",
			Usage:       "Value for key",
			Destination: &Args.Value,
		},
		cli.BoolFlag{
			Name:        "encryption",
			Usage:       "Set Encryption For Value",
			Destination: &Args.Encryption,
		},
	}
	command.Action = func(c *cli.Context) error {
		client, err := handler.Lib.GetClients(Args.EnvPath)
		if Args.TemplatePath != "" {
			log.Println("PUT BY TEMPLATE")
			return nil
		}

		params := client.GetParameters()
		params.Path = Args.Key
		if Args.Encryption {
			encValue, _ := handler.Lib.EncryptValue(Args.Value)
			params.Value = string(encValue)
		} else {
			params.Value = Args.Value
		}

		client.Put(params)
		result, err := client.Get(params.Path)
		log.Println("Create Revision: ", result.Kvs[0].CreateRevision)
		log.Println("Mod Revision: ", result.Kvs[0].ModRevision)
		return err
	}

	return command
}
