package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

var Version = "0.0.4"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize cnats with configuration file",
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
			Name:        "spub",
			Usage:       "cnats spub subject 'json_payload'",
			Description: "Publish message to streaming nats with json payload",
			Action:      spubActionHandler,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "client",
					Usage: "If provided overrides default client id.",
				},
			},
		},
		{
			Name:        "pub",
			Usage:       "cnats pub subject 'json_payload'",
			Description: "Publish message to nats with json payload",
			Action:      pubActionHandler,
			Flags:       []cli.Flag{},
		},
		{
			Name:        "ssub",
			Usage:       "Subscribe streaming nats to messages. Eg. cnats ssub subject1 subject2",
			Description: "If subjects are not provided cnats subscribes to all known subjects.",
			Action:      ssubActionHandler,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "client",
					Usage: "If provided overrides default client id.",
				},
			},
		},
		{
			Name:        "sub",
			Usage:       "Subscribe to nats messages. Eg. cnats sub subject1 subject2",
			Description: "If subjects are not provided cnats subscribes to all known subjects.",
			Action:      subActionHandler,
			Flags:       []cli.Flag{},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
