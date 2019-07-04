package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func createStanConn(conf *config) (stan.Conn, error) {
	conn, err := stan.Connect(
		conf.ClusterID,
		conf.ClientID,
		stan.NatsURL(conf.URL),
	)

	if err != nil {
		return nil, fmt.Errorf("could not create stan conn: %v", err)
	}

	return conn, nil
}

func createConn(conf *config) (*nats.Conn, error) {
	conn, err := nats.Connect(conf.URL)

	if err != nil {
		return nil, fmt.Errorf("could not create stan conn: %v", err)
	}

	return conn, nil
}
