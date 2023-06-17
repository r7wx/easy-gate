package config

import (
	"fmt"
	"os"
)

// GetConfigPath - Get the path to the configuration file
func GetConfigPath(args []string) (string, error) {
	cfgFilePath := os.Getenv(cfgPathEnv)
	if cfgFilePath != "" {
		return cfgFilePath, nil
	}

	if len(args) <= 1 {
		return "", fmt.Errorf("no configuration file provided")
	}

	return args[1], nil
}
