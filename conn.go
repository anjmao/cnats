package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func createConn(conf *config) (stan.Conn, error) {
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
