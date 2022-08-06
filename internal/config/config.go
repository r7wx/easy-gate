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
	"encoding/json"

	"github.com/r7wx/easy-gate/internal/errors"
	"github.com/r7wx/easy-gate/internal/models"
	"gopkg.in/yaml.v3"
)

// Service - Easy Gate service configuration struct
type Service struct {
	Icon        string   `json:"icon" yaml:"icon"`
	Name        string   `json:"name" yaml:"name"`
	URL         string   `json:"url" yaml:"url"`
	Groups      []string `json:"groups" yaml:"groups"`
	HealthCheck bool     `json:"health_check" yaml:"health_check"`
}

// Note - Easy Gate note configuration struct
type Note struct {
	Name   string   `json:"name" yaml:"name"`
	Text   string   `json:"text" yaml:"text"`
	Groups []string `json:"groups" yaml:"groups"`
}

// Theme - Easy Gate theme configuration struct
type Theme struct {
	Background string `json:"background" yaml:"background"`
	Foreground string `json:"foreground" yaml:"foreground"`
}

// Config - Easy Gate configuration struct
type Config struct {
	Theme       Theme          `json:"theme" yaml:"theme"`
	Addr        string         `json:"addr" yaml:"addr"`
	Title       string         `json:"title" yaml:"title"`
	CertFile    string         `json:"cert_file" yaml:"cert_file"`
	KeyFile     string         `json:"key_file" yaml:"key_file"`
	Groups      []models.Group `json:"groups" yaml:"groups"`
	Services    []Service      `json:"services" yaml:"services"`
	Notes       []Note         `json:"notes" yaml:"notes"`
	BehindProxy bool           `json:"behind_proxy" yaml:"behind_proxy"`
	UseTLS      bool           `json:"use_tls" yaml:"use_tls"`
}

type format int

const (
	jsonFormat format = iota + 1
	yamlFormat
)

// Unmarshal - Unmarshal config bytes into config struct
func Unmarshal(configBytes []byte) (*Config, error) {
	config := &Config{}

	allowedFormats := []format{jsonFormat, yamlFormat}
	for _, format := range allowedFormats {
		switch format {
		case jsonFormat:
			err := json.Unmarshal(configBytes, config)
			if err == nil {
				return config, nil
			}
		case yamlFormat:
			err := yaml.Unmarshal(configBytes, config)
			if err == nil {
				return config, nil
			}
		}
	}

	return nil, errors.NewEasyGateError(
		errors.InvalidFormat,
		errors.ConfigurationFile, "",
	)
}
