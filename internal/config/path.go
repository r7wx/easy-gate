package config

import (
	"fmt"
	"net/url"
	"os"
)

const (
	cfgPathEnv  = "EASY_GATE_CONFIG_PATH"
	rootPathEnv = "EASY_GATE_ROOT_PATH"
)

// GetRootPath - Get root directory path from env
// or return the default value "/"
func GetRootPath() string {
	rootPath := os.Getenv(rootPathEnv)
	return JoinUrlPath("/", rootPath)
}

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

// JoinUrlPath - Wrapper around url.JoinPath
func JoinUrlPath(base string, elem ...string) string {
	path, err := url.JoinPath(base, elem...)
	if err != nil {
		url.JoinPath("/", elem...)
	}
	return path
}
