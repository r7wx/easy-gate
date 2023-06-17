package config

import (
	"fmt"
	"regexp"
)

func isHexColor(color string) bool {
	if len(color) < 4 || len(color) > 7 {
		return false
	}

	if color[0] != '#' {
		return false
	}

	for i := 1; i < len(color); i++ {
		c := color[i]
		if (c >= '0' && c <= '9') || (c >= 'a' &&
			c <= 'f') || (c >= 'A' && c <= 'F') {
			continue
		}
		return false
	}

	return true
}

func isURL(url string) bool {
	r, _ := regexp.Compile(
		`^(https?|ftp)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	return r.MatchString(url)
}

func validateConfig(cfg *Config) error {
	if !isHexColor(cfg.Theme.Background) {
		return fmt.Errorf("invalid background color")
	}
	if !isHexColor(cfg.Theme.Foreground) {
		return fmt.Errorf("invalid foreground color")
	}

	for _, service := range cfg.Services {
		if !isURL(service.URL) {
			return fmt.Errorf("invalid URL for service %s", service.Name)
		}
	}

	return nil
}
