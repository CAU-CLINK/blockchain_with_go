package main

import (
	"log"
	"os"
	"time"

	"github.com/CAU-CLINK/blockchain_with_go/common"
	"github.com/CAU-CLINK/blockchain_with_go/conf"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "clink-education-chain-core"
	app.Version = "0.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "KSJ",
			Email: "leesd556@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
			Usage: "name for config",
		},
	}
	app.Commands = []cli.Command{}

	app.Before = func(c *cli.Context) error {
		if configPath := c.String("config"); configPath != "" {
			absPath, err := common.RelativeToAbsolutePath(configPath)
			if err != nil {
				return err
			}
			conf.ConfigPath(absPath)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
