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
	"io/ioutil"
	"os"
	"testing"

	"github.com/r7wx/easy-gate/internal/share"
)

var sampleJSONConfig string = `{
"addr": "0.0.0.0:8080",
"use_tls": false,
"cert_file": "",
"key_file": "",
"behind_proxy": false,
"title": "Easy Gate",
"theme": {
	"background": "#FFFFFF",
	"foreground": "#000000"
},
"groups": [],
"services": [],
"notes": []
}`

var sampleYMLConfig string = `addr: 0.0.0.0:8080
use_tls: false
cert_file: ''
key_file: ''
behind_proxy: false
title: Easy Gate
theme:
  background: '#FFFFFF'
  foreground: '#000000'
groups: []
services: []
notes: []`

func TestUnmarshal(t *testing.T) {
	_, err := Unmarshal([]byte(sampleJSONConfig))
	if err != nil {
		t.Fatal(err)
	}
	_, err = Unmarshal([]byte(sampleYMLConfig))
	if err != nil {
		t.Fatal(err)
	}
	_, err = Unmarshal([]byte("X"))
	if err == nil {
		t.Fatal()
	}
}

func TestLoadEnv(t *testing.T) {
	os.Setenv(share.CFGEnv, sampleJSONConfig)
	_, _, err := Load("")
	if err != nil {
		t.Fatal(err)
	}
	os.Unsetenv(share.CFGEnv)
}

func TestLoadFile(t *testing.T) {
	cfgFile, err := ioutil.TempFile(".", "easy_gate_test_")
	if err != nil {
		t.Fatal(err)
	}
	cfgFile.WriteString(sampleJSONConfig)
	_, _, err = Load(cfgFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(cfgFile.Name())

	cfgFile, err = ioutil.TempFile(".", "easy_gate_test_")
	if err != nil {
		t.Fatal(err)
	}
	cfgFile.WriteString("XXXX")
	_, _, err = Load(cfgFile.Name())
	if err == nil {
		t.Fatal()
	}
	defer os.Remove(cfgFile.Name())

	cfgFile, err = ioutil.TempFile(".", "easy_gate_test_")
	if err != nil {
		t.Fatal(err)
	}
	wrongConfig := `{
"addr": "0.0.0.0:8080",
"use_tls": false,
"cert_file": "",
"key_file": "",
"behind_proxy": false,
"title": "Easy Gate",
"theme": {
	"background": "TEST123",
	"foreground": "TEST"
},
"groups": [],
"services": [],
"notes": []
}`
	cfgFile.WriteString(wrongConfig)
	_, _, err = Load(cfgFile.Name())
	if err == nil {
		t.Fatal()
	}
	defer os.Remove(cfgFile.Name())

	_, _, err = Load("H()2NOTEXISTENT.NOT")
	if err == nil {
		t.Fatal()
	}
}

func TestGetPath(t *testing.T) {
	os.Setenv(share.CFGPathEnv, "test.json")
	_, err := GetConfigPath([]string{})
	if err != nil {
		t.Fatal(err)
	}
	os.Unsetenv(share.CFGPathEnv)

	_, err = GetConfigPath([]string{})
	if err == nil {
		t.Fatal()
	}

	path, err := GetConfigPath([]string{"", "test.json"})
	if err != nil {
		t.Fatal()
	}
	if path != "test.json" {
		t.Fatal()
	}
}

func TestValidate(t *testing.T) {
	cfg := Config{}
	err := validateConfig(&cfg)
	if err == nil {
		t.Fatal()
	}

	cfg.Theme.Background = "#FFF"
	err = validateConfig(&cfg)
	if err == nil {
		t.Fatal()
	}

	cfg.Theme.Foreground = "#____"
	err = validateConfig(&cfg)
	if err == nil {
		t.Fatal()
	}

	cfg.Theme.Foreground = "#FFF"
	cfg.Services = []Service{
		{URL: "javascript:alert(1)"},
	}
	err = validateConfig(&cfg)
	if err == nil {
		t.Fatal()
	}

	cfg.Services = []Service{
		{URL: "https://www.google.com"},
	}
	err = validateConfig(&cfg)
	if err != nil {
		t.Fatal()
	}
}
