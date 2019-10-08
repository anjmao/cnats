package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

var (
	errPubSubjectRequired = errors.New("pub: subject is required")
	errPubPayloadRequired = errors.New("pub: payload is required")
)

func spubActionHandler(c *cli.Context) error {
	subject := c.Args().Get(0)
	if subject == "" {
		return errPubSubjectRequired
	}

	payload := c.Args().Get(1)
	if payload == "" {
		return errPubPayloadRequired
	}

	conf, err := readConfig()
	if err != nil {
		return err
	}

	clientIDOverride := c.String("client")
	if clientIDOverride != "" {
		conf.ClientID = clientIDOverride
	}

	conn, err := createStanConn(conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Publish(subject, []byte(payload)); err != nil {
		return fmt.Errorf("pub: could not publish message: %v", err)
	}
	return nil
}

func pubActionHandler(c *cli.Context) error {
	subject := c.Args().Get(0)
	if subject == "" {
		return errPubSubjectRequired
	}

	payload := c.Args().Get(1)
	if payload == "" {
		return errPubPayloadRequired
	}

	conf, err := readConfig()
	if err != nil {
		return err
	}

	clientIDOverride := c.String("client")
	if clientIDOverride != "" {
		conf.ClientID = clientIDOverride
	}

	conn, err := createConn(conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Publish(subject, []byte(payload)); err != nil {
		return fmt.Errorf("pub: could not publish message: %v", err)
	}
	return nil
}
