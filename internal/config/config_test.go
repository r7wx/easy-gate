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
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

var testCfg Config = Config{
	Addr:        ":8080",
	UseTLS:      false,
	CertFile:    "",
	KeyFile:     "",
	BehindProxy: false,
	Title:       "Test",
	Icon:        "fa-solid fa-cubes",
	Motd:        "",
	Theme: Theme{
		Background: "#ffffff",
		Foreground: "#000000",
	},
	Groups: []Group{
		{
			Name:   "Group 1",
			Subnet: "127.0.0.1/32",
		},
	},
	Services: []Service{
		{
			Name:   "Service 1",
			Icon:   "fa-solid fa-cube",
			URL:    "http://test:8080",
			Groups: []string{},
		},
		{
			Name:   "Service 2",
			Icon:   "fa-solid fa-cube",
			URL:    "http://test2:8080",
			Groups: []string{"Group 1"},
		},
	},
	Notes: []Note{
		{
			Name:   "Note 1",
			Text:   "This is a test note",
			Groups: []string{},
		},
		{
			Name:   "Note 2",
			Text:   "This is another test note",
			Groups: []string{"Group 1"},
		},
	},
}

const (
	testJSONPath = "./test-config.json"
	testYAMLPath = "./test-config.yml"
)

func TestMain(m *testing.M) {
	cfgJSON, err := json.Marshal(testCfg)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(testJSONPath,
		cfgJSON, 0644)
	if err != nil {
		panic(err)
	}

	cfgYAML, err := yaml.Marshal(testCfg)
	err = ioutil.WriteFile(testYAMLPath,
		cfgYAML, 0644)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Remove(testJSONPath)
	os.Remove(testYAMLPath)
	os.Exit(exitCode)
}

func TestJSON(t *testing.T) {
	routine := NewRoutine(testJSONPath, 1*time.Second)
	go routine.Start()

	cfg, err := routine.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*cfg, testCfg) {
		t.Fatal("JSON configuration not parsed correctly")
	}
}

func TestYAML(t *testing.T) {
	routine := NewRoutine(testYAMLPath, 1*time.Second)
	go routine.Start()

	cfg, err := routine.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*cfg, testCfg) {
		t.Fatal("YAML configuration not parsed correctly")
	}
}

func TestFailConditions(t *testing.T) {
	_, err := Unmarshal([]byte("AAAAAAA"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	routine := Routine{
		FilePath: "TEST_FILE_DOES_NOT_EXIST.NOT_EXIST",
		Interval: 1 * time.Second,
	}
	go routine.Start()
	time.Sleep(2 * time.Second)

	_, err = routine.GetConfiguration()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestHexColors(t *testing.T) {
	if !isHexColor("#ff0000") {
		t.Fatal("Expected true, got false")
	}
	if !isHexColor("#f00") {
		t.Fatal("Expected true, got false")
	}
	if !isHexColor("#ffff") {
		t.Fatal("Expected true, got false")
	}
	if isHexColor("FFFFFF") {
		t.Fatal("Expected false, got true")
	}
	if isHexColor("#FFFFFFF") {
		t.Fatal("Expected false, got true")
	}
	if isHexColor("#") {
		t.Fatal("Expected false, got true")
	}
	if isHexColor("#FFG") {
		t.Fatal("Expected false, got true")
	}
	if isHexColor("32984327493827@@@AA") {
		t.Fatal("Expected false, got true")
	}
}

func TestURLs(t *testing.T) {
	if !isURL("http://example.com") {
		t.Fatal("Expected true, got false")
	}
	if !isURL("https://example.com") {
		t.Fatal("Expected true, got false")
	}
	if !isURL("https://example.com/test/test.xy") {
		t.Fatal("Expected true, got false")
	}
	if !isURL("https://example.com/test/test.xy?test=test") {
		t.Fatal("Expected true, got false")
	}
	if !isURL("https://example.com/test/test.xy?test=test#test") {
		t.Fatal("Expected true, got false")
	}
	if !isURL("ftp://example.com") {
		t.Fatal("Expected true, got false")
	}
	if isURL("example.internal.priv") {
		t.Fatal("Expected false, got true")
	}
	if isURL("test.test") {
		t.Fatal("Expected false, got true")
	}
	if isURL("example") {
		t.Fatal("Expected false, got true")
	}
	if isURL("javascript:void(0)") {
		t.Fatal("Expected false, got true")
	}
	if isURL("javascript:alert(1)") {
		t.Fatal("Expected false, got true")
	}
	if isURL("javascript: alert(1)") {
		t.Fatal("Expected false, got true")
	}
}

func TestIcons(t *testing.T) {
	if !isIcon("fa-brands fa-github") {
		t.Fatal("Expected true, got false")
	}
	if !isIcon("fa-regular fa-cube") {
		t.Fatal("Expected true, got false")
	}
	if !isIcon("fa-solid fa-flask-vial") {
		t.Fatal("Expected true, got false")
	}
	if isIcon("") {
		t.Fatal("Expected false, got true")
	}
	if isIcon("bg-white text-red") {
		t.Fatal("Expected false, got true")
	}
	if isIcon("fa-brands fa-github fa-brands fa-github") {
		t.Fatal("Expected false, got true")
	}
}
