package config

import (
	"os"
	"pr9/internal/delivery"
	"pr9/pkg/mongoconnector"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mongo  mongoconnector.Config `yaml:"Mongo"`
	Server delivery.Config       `yaml:"Server"`
}

var configPath = "./config/config.yaml"

func LoadConfig() (config *Config, err error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, err
}
