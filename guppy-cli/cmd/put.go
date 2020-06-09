package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func (handler *CLIMapping) put() cli.Command {
	command := cli.Command{}
	command.Name = "put"
	command.Usage = "put [option]"
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
			handler.putByPath(Args.EnvPath)
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

		_, err = client.Put(params)
		if err != nil {
			return err
		}
		result, err := client.Get(params.Path)
		if err != nil {
			return err
		}
		fmt.Println("Paramater Insert")
		headers := []string{
			"Path",
			"Value",
			"Create Revision",
			"Mod Revision",
		}
		handler.Utils.GenerateGetTable(headers, result.Kvs)
		return nil
	}

	return command
}

func (handler *CLIMapping) putByPath(envPath string) error {
	path, err := handler.Utils.CheckTemplateFile(Args.TemplatePath)
	if err != nil {
		return err
	}

	objectData, err := handler.Utils.ParsingYAML(path)
	if err != nil {
		return err
	}
	handler.Lib.PutByPath(envPath, objectData)
	return nil
}
