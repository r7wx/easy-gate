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
	"fmt"
	"regexp"

	"github.com/r7wx/easy-gate/internal/errors"
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
			fmt.Println(color)
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
		return errors.NewEasyGateError(
			errors.InvalidColor,
			errors.Root,
			"background",
		)
	}
	if !isHexColor(cfg.Theme.Foreground) {
		return errors.NewEasyGateError(
			errors.InvalidColor,
			errors.Root,
			"foreground",
		)
	}

	for _, service := range cfg.Services {
		if !isURL(service.URL) {
			return errors.NewEasyGateError(
				errors.InvalidURL,
				errors.Service,
				service.Name,
			)
		}
	}

	return nil
}
