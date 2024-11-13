package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	CFG_FILE = "GATEWAY_CONFIG"
)

type HttpCfg struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	HttpConfigs HttpCfg `json:"http"`
}

func LoadConfigFile(file string) *Config {
	data, err := os.ReadFile(file)
	if err != nil {
		msg := fmt.Sprintf("Cannot find config file %s: %v", file, err)
		log.Fatal(msg)
	}

	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		msg := fmt.Sprintf("Failed to load config: %v", err)
		log.Fatal(msg)
	}

	return &c
}
