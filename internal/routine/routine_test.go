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

package routine

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/models"
)

var cfgFilePath string

func TestMain(m *testing.M) {
	configContent := `{
"addr": "127.0.0.1:8080",
"use_tls": false,
"cert_file": "",
"key_file": "",
"behind_proxy": false,
"title": "Easy Gate",
"theme": {
	"background": "#FFFFFF",
	"foreground": "#000000",
	"health_ok": "#FFFFFF",
	"health_bad": "#000000",
	"health_inactive": "#FFFFFF",
},
"groups": [],
"services": [],
"notes": []
}`
	cfgFile, err := ioutil.TempFile(".", "easy_gate_test_")
	if err != nil {
		log.Fatal("Unable to write tmp files for test")
	}
	cfgFile.WriteString(configContent)
	cfgFilePath = cfgFile.Name()

	exitCode := m.Run()

	os.Remove(cfgFilePath)

	os.Exit(exitCode)
}

func TestRoutine(t *testing.T) {
	_, err := NewRoutine("", 1*time.Millisecond)
	if err == nil {
		t.Fatal()
	}

	testRoutine, err := NewRoutine(cfgFilePath, 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	_, err = testRoutine.GetStatus()
	if err != nil {
		t.Fatal()
	}
}

func TestHealth(t *testing.T) {
	service := config.Service{
		URL:         "https://www.google.com",
		HealthCheck: true,
	}
	status := checkHealth(service)
	if status == models.HealthUndefined {
		t.Fatal()
	}

	service = config.Service{
		URL:         "https://www.google.com",
		HealthCheck: false,
	}
	status = checkHealth(service)
	if status != models.HealthUndefined {
		t.Fatal()
	}
}

func TestGetServices(t *testing.T) {
	cfg := config.Config{
		Services: []config.Service{
			{
				Icon:        "",
				Name:        "Test 1",
				URL:         "https://test.test",
				HealthCheck: true,
			},
		},
	}

	services := getServices(&cfg)
	if services[0].Name != "Test 1" {
		t.Fatal()
	}
}

func TestGetNotes(t *testing.T) {
	cfg := config.Config{
		Notes: []config.Note{
			{
				Name: "Test 1",
				Text: "...",
			},
		},
	}

	notes := getNotes(&cfg)
	if notes[0].Name != "Test 1" {
		t.Fatal()
	}
}
