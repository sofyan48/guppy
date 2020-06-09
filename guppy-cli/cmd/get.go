package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"go.etcd.io/etcd/clientv3"
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
			Required:    true,
		},
		cli.BoolFlag{
			Name:        "with-decryption",
			Usage:       "Set Decryption For Value",
			Destination: &Args.Encryption,
		},
	}
	command.Action = func(c *cli.Context) error {
		result := &clientv3.GetResponse{}
		client, err := handler.Lib.GetClients(Args.EnvPath)
		if err != nil {
			return err
		}

		headers := []string{
			"Path",
			"Value",
			"Create Revision",
			"Mod Revision",
		}

		if len(c.Args()) <= 0 {
			fmt.Println("Select Get subcommand items or path")
			return nil
		}
		switch c.Args()[0] {
		case "items":
			result, err = client.Get(Args.Key)
			if err != nil {
				return err
			}
		case "path":
			result, err = client.GetByPath(Args.Key)
			if err != nil {
				return err
			}
		}
		fmt.Println("Result: ")
		handler.Utils.GenerateGetTable(headers, result.Kvs)
		return nil
	}

	return command
}
