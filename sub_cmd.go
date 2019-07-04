package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/urfave/cli"
)

func ssubActionHandler(c *cli.Context) error {
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

	subjects := c.Args()
	if len(subjects) == 0 {
		channels, err := getStreamingChannels(conf.URL)
		if err != nil {
			return fmt.Errorf("could not get channels from monitoring api: %v", err)
		}
		subjects = channels.Names
	}

	for _, sub := range subjects {
		go handleStanSubscription(conn, sub)
	}

	var stop string
	fmt.Scanln(&stop)

	return nil
}

func handleStanSubscription(conn stan.Conn, subject string) {
	fmt.Printf("subscribing to \"%s\"\n", subject)
	_, err := conn.Subscribe(subject, func(msg *stan.Msg) {
		fmt.Println("---------------------------------------")
		fmt.Printf("subject:   %s\n", msg.Subject)
		fmt.Printf("timestamp: %d\n", msg.Timestamp)
		fmt.Printf("payload:   %s\n", string(msg.Data))

	})
	if err != nil {
		fmt.Printf("could not subscribe to %s: %v", subject, err)
	}
}

func subActionHandler(c *cli.Context) error {
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

	subjects := c.Args()
	if len(subjects) == 0 {
		return nil
	}

	for _, sub := range subjects {
		go handleSubscription(conn, sub)
	}

	var stop string
	fmt.Scanln(&stop)

	return nil
}

func handleSubscription(conn *nats.Conn, subject string) {
	fmt.Printf("subscribing to \"%s\"\n", subject)
	_, err := conn.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("---------------------------------------")
		fmt.Printf("subject:   %s\n", msg.Subject)
		fmt.Printf("payload:   %s\n", string(msg.Data))

	})
	if err != nil {
		fmt.Printf("could not subscribe to %s: %v", subject, err)
	}
}

type channelsz struct {
	ClusterID string   `json:"cluster_id"`
	Names     []string `json:"names"`
}

func getStreamingChannels(natsURL string) (*channelsz, error) {
	parts := strings.Split(natsURL, ":")
	host := parts[1]
	stanMonitoringURL := fmt.Sprintf("http:%s:8222/streaming", host)
	url := fmt.Sprintf("%s/channelsz", stanMonitoringURL)
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	channels := &channelsz{}
	if err := json.NewDecoder(res.Body).Decode(channels); err != nil {
		return nil, err
	}

	return channels, nil
}
