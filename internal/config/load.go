package config

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

const cfgEnv = "EASY_GATE_CONFIG"

// Load - Load configuration from environment or file
func Load(filePath string) (*Config, string, error) {
	envCfg := os.Getenv(cfgEnv)
	if envCfg != "" {
		return loadConfig([]byte(envCfg))
	}

	cfgFile, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer cfgFile.Close()

	fileData, _ := io.ReadAll(cfgFile)
	return loadConfig(fileData)
}

func loadConfig(cfgData []byte) (*Config, string, error) {
	checksum := checksum(cfgData)
	cfg, err := Unmarshal(cfgData)
	if err != nil {
		return nil, "", err
	}

	if err := validateConfig(cfg); err != nil {
		return nil, "", err
	}

	return cfg, checksum, nil
}

func checksum(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
