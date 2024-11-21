package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Filters []string `yaml:"filters"`
}

func (c *Config) LoadConfig(filepath string) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
