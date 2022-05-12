/*
MIT License

Copyright (c) 2022 r7wx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package config

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"

	"github.com/r7wx/easy-gate/internal/share"
)

// LoadConfig - Load configuration from environment or file
func LoadConfig(filePath string) (*Config, string, error) {
	envCfg := os.Getenv(share.CFGEnv)
	if envCfg != "" {
		log.Println("Loading configuration from environment")
		return loadConfig([]byte(envCfg))
	}

	cfgFile, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer cfgFile.Close()

	fileData, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		return nil, "", err
	}

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
