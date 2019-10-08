package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	errInitClusterRequired = errors.New("init: cluster id is required")
	errInitURLRequired     = errors.New("init: connection url is required")
)

func initActionHandler(c *cli.Context) error {
	clusterID := c.String("cluster")
	if clusterID == "" {
		return errInitClusterRequired
	}
	clientID := c.String("client")
	if clientID == "" {
		clientID = generateClientID()
	}
	url := c.String("url")
	if url == "" {
		return errInitURLRequired
	}

	conf := &config{
		ClusterID: clusterID,
		ClientID:  clientID,
		URL:       url,
	}
	return saveConfig(conf)
}

func generateClientID() string {
	b := make([]byte, 5)
	_, err := rand.Read(b)
	if err != nil {
		return "cnat"
	}

	bhex := hex.EncodeToString(b)
	return fmt.Sprintf("cnat_%s", bhex)
}
