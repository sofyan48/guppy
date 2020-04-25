package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sofyan48/guppy/guppy"
	"github.com/sofyan48/guppy/guppy/config"
)

func main() {
	config := config.NewConfig()
	config.DialTimeOut = 5
	config.Urls = []string{"localhost:32770"}
	client, err := guppy.Client(config).New()
	if err != nil {
		log.Println("Client Not Connected: ", err)
		os.Exit(1)
	}
	params := client.GetParameters()
	params.Path = "/prd/general/service/database/DATABASE_HOST"
	params.Value = "localhost"
	result, err := client.Put(params)
	if err != nil {
		log.Println("Put Error: ", err)
		os.Exit(1)
	}
	fmt.Println(result)

}
