package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Filters    []string `yaml:"filters"`
	Iterations int      `yaml:"iterations"`
}

func getDefaultConfigPath() (defaultPath string, err error) {
	programName := "uptfs"
	configFileName := "config.yaml"

	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, programName, configFileName), nil
	}

	if home := os.Getenv("HOME"); home != "" {
		return filepath.Join(home, programName, configFileName), nil
	}

	return "", errors.New("both XDG_CONFIG_HOME and HOME are not set, can't proceed")
}

func (c *Config) LoadConfig(filepath string) *Config {
	if filepath == "" {
		var err error
		filepath, err = getDefaultConfigPath()

		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}

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

	return c
}
