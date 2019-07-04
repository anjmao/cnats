package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

var Version = "0.0.3"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize cnat with configuration file",
			Action: initActionHandler,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Streaming NATS Cluster ID.",
					Value: "test-cluster",
				},
				cli.StringFlag{
					Name:  "client",
					Usage: "Streaming NATS ID. Generates random client id if not provided.",
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "Connection url.",
					Value: "nats://localhost:4222",
				},
			},
		},
		{
			Name:        "pub",
			Usage:       "cnat subject 'json_payload'",
			Description: "Publish message with json payload",
			Action:      pubActionHandler,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "client",
					Usage: "If provided overrides default client id.",
				},
			},
		},
		{
			Name:        "sub",
			Usage:       "Subscribe to messages. Eg. cnat subject1 subject2",
			Description: "If subjects are not provided cnat subscribes to all known subjects.",
			Action:      subActionHandler,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "client",
					Usage: "If provided overrides default client id.",
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
