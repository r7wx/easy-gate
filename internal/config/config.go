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

// Group - Self Gate group configuration struct
type Group struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
}

// Service - Self Gate service configuration struct
type Service struct {
	Icon   string   `json:"icon"`
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Groups []string `json:"groups"`
}

// Note - Self Gate note configuration struct
type Note struct {
	Name   string   `json:"name"`
	Text   string   `json:"text"`
	Groups []string `json:"groups"`
}

// Theme - Self Gate theme configuration struct
type Theme struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
}

// Config - Self Gate configuration struct
type Config struct {
	Theme       Theme     `json:"theme"`
	Addr        string    `json:"addr"`
	Title       string    `json:"title"`
	CertFile    string    `json:"cert_file"`
	KeyFile     string    `json:"key_file"`
	Icon        string    `json:"icon"`
	Motd        string    `json:"motd"`
	Groups      []Group   `json:"groups"`
	Services    []Service `json:"services"`
	Notes       []Note    `json:"notes"`
	BehindProxy bool      `json:"behind_proxy"`
	UseTLS      bool      `json:"use_tls"`
}
