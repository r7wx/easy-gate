package config

import (
	"os"
	"testing"
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
	os.Setenv(cfgEnv, sampleJSONConfig)
	_, _, err := Load("")
	if err != nil {
		t.Fatal(err)
	}
	os.Unsetenv(cfgEnv)
}

func TestLoadFile(t *testing.T) {
	cfgFile, err := os.CreateTemp(".", "easy_gate_test_")
	if err != nil {
		t.Fatal(err)
	}
	cfgFile.WriteString(sampleJSONConfig)
	_, _, err = Load(cfgFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(cfgFile.Name())

	cfgFile, err = os.CreateTemp(".", "easy_gate_test_")
	if err != nil {
		t.Fatal(err)
	}
	cfgFile.WriteString("XXXX")
	_, _, err = Load(cfgFile.Name())
	if err == nil {
		t.Fatal()
	}
	defer os.Remove(cfgFile.Name())

	cfgFile, err = os.CreateTemp(".", "easy_gate_test_")
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
	os.Setenv(cfgPathEnv, "test.json")
	_, err := GetConfigPath([]string{})
	if err != nil {
		t.Fatal(err)
	}
	os.Unsetenv(cfgPathEnv)

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
		{URL: "XXX"},
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
