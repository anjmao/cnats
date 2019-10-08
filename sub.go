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
		channels, err := getStanSubscribers(conf.URL)
		if err != nil {
			return fmt.Errorf("could not get subscribers from monitoring api: %v", err)
		}
		subjects = channels
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

	conn, err := createConn(conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	subjects := c.Args()
	if len(subjects) == 0 {
		subs, err := getNatsSubscribers(conf.URL)
		if err != nil {
			return fmt.Errorf("could not get subscribers from monitoring api: %v", err)
		}
		subjects = subs
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

type stanSubsz struct {
	ClusterID string   `json:"cluster_id"`
	Names     []string `json:"names"`
}

func getStanSubscribers(natsURL string) ([]string, error) {
	parts := strings.Split(natsURL, ":")
	host := parts[1]
	monitoringURL := fmt.Sprintf("http:%s:8222/streaming", host)
	url := fmt.Sprintf("%s/channelsz", monitoringURL)
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	subs := &stanSubsz{}
	if err := json.NewDecoder(res.Body).Decode(subs); err != nil {
		return nil, err
	}

	return subs.Names, nil
}

func getNatsSubscribers(natsURL string) ([]string, error) {
	parts := strings.Split(natsURL, ":")
	host := parts[1]
	monitoringURL := fmt.Sprintf("http:%s:8222", host)
	url := fmt.Sprintf("%s/subsz?subs=1", monitoringURL)
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	subsz := &natsSubsz{}
	if err := json.NewDecoder(res.Body).Decode(subsz); err != nil {
		return nil, err
	}

	var subs []string
	for _, s := range subsz.List {
		subs = append(subs, s.Subject)
	}
	return subs, nil
}

type natsSubscription struct {
	Subject string `json:"subject"`
}

type natsSubsz struct {
	List []natsSubscription `json:"subscriptions_list"`
}
