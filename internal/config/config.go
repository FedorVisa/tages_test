package config

import (
	"gopkg.in/yaml.v2"

	"log"
	"os"
)

type Config struct {
	Port     string `yaml:"port"`
	Filepath string `yaml:"path"`
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error in reading of configuration file %v", err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error in unmarshaling the file: %v", err)
	}
	return &config, nil
}
