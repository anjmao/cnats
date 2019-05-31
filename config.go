package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type config struct {
	ClusterID string `json:"cluster_id"`
	ClientID  string `json:"client_id"`
	URL       string `json:"url"`
}

const (
	configDirName  = ".cnat"
	configFileName = "config.json"
)

func saveConfig(c *config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filePath := path.Join(homeDir, configDirName)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(filePath, 0744); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("could not stat dir: %v", err)
		}
	}

	fileName := path.Join(filePath, configFileName)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return fmt.Errorf("could not create config file: %v", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(c); err != nil {
		return fmt.Errorf("coult not write to file: %v", err)
	}

	return nil
}

func readConfig() (*config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fileName := path.Join(homeDir, configDirName, configFileName)

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %v", err)
	}

	res := &config{}
	if err := json.Unmarshal(bytes, res); err != nil {
		return nil, fmt.Errorf("could not unmarshal json to config: %v", err)
	}
	return res, nil
}
