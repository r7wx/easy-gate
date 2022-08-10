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

package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/models"
	"github.com/r7wx/easy-gate/internal/routine"
)

var testConfigFilePath string

func TestMain(m *testing.M) {
	testCfg := config.Config{
		Addr:        "127.0.0.1:8080",
		UseTLS:      false,
		CertFile:    "",
		KeyFile:     "",
		BehindProxy: false,
		Title:       "Test",
		Theme: models.Theme{
			Background: "#ffffff",
			Foreground: "#000000",
		},
		Groups: []models.Group{
			{
				Name:   "test",
				Subnet: "192.168.1.1/24",
			},
			{
				Name:   "test2",
				Subnet: "10.1.5.1/24",
			},
		},
		Services: []config.Service{
			{
				Name:   "service1",
				URL:    "http://example.com/service1",
				Groups: []string{},
			},
			{
				Name:   "service2",
				URL:    "http://example.com/service2",
				Groups: []string{"test"},
			},
			{
				Name:   "service3",
				URL:    "http://example.com/service3",
				Groups: []string{"test2"},
			},
		},
		Notes: []config.Note{
			{
				Name:   "note1",
				Groups: []string{},
			},
			{
				Name:   "note2",
				Groups: []string{"test"},
			},
			{
				Name:   "note3",
				Groups: []string{"test2"},
			},
		},
	}

	cfgJSON, err := json.Marshal(testCfg)
	if err != nil {
		log.Fatal(err)
	}
	cfgFile, err := ioutil.TempFile(".", "easy_gate_test_")
	if err != nil {
		log.Fatal("Unable to write tmp files for test")
	}
	if err != nil {
		log.Fatal(err)
	}

	testConfigFilePath = cfgFile.Name()
	cfgFile.Write(cfgJSON)

	exitCode := m.Run()

	os.Remove(testConfigFilePath)
	os.Exit(exitCode)
}

func TestService(t *testing.T) {
	routine, err := routine.NewRoutine(testConfigFilePath, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	go routine.Start()
	service := NewService(routine)

	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
	w := httptest.NewRecorder()
	service.data(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}
	service.webFS(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	routine.Lock()
	routine.Status.BehindProxy = true
	routine.Unlock()

	req = httptest.NewRequest(http.MethodGet, "/api/data", nil)
	w = httptest.NewRecorder()
	service.data(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal()
	}
	service.webFS(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal()
	}

	req = httptest.NewRequest(http.MethodGet, "/api/data", nil)
	req.Header.Set("X-Forwarded-For", "127.0.0.1")
	w = httptest.NewRecorder()
	service.data(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}
	service.webFS(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	routine.Lock()
	routine.Status.BehindProxy = false
	routine.Error = fmt.Errorf("Test error")
	routine.Unlock()

	req = httptest.NewRequest(http.MethodGet, "/api/data", nil)
	w = httptest.NewRecorder()
	service.data(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	routine.Lock()
	routine.Error = nil
	routine.Unlock()
}

func TestGetServices(t *testing.T) {
	routine, err := routine.NewRoutine(testConfigFilePath,
		8*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	go routine.Start()

	service := NewService(routine)
	cfg, _ := service.Routine.GetStatus()

	services := service.getServices(cfg, "192.168.1.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service2" {
			t.Fail()
		}
	}

	services = service.getServices(cfg, "10.1.5.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service3" {
			t.Fail()
		}
	}

	services = service.getServices(cfg, "1.1.1.1")
	for _, s := range services {
		if s.Name != "service1" {
			t.Fail()
		}
	}
}

func TestGetNotes(t *testing.T) {
	routine, err := routine.NewRoutine(testConfigFilePath,
		8*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	go routine.Start()

	service := NewService(routine)
	cfg, _ := service.Routine.GetStatus()

	notes := service.getNotes(cfg, "192.168.1.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note2" {
			t.Fail()
		}
	}

	notes = service.getNotes(cfg, "10.1.5.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note3" {
			t.Fail()
		}
	}

	notes = service.getNotes(cfg, "1.1.1.1")
	for _, n := range notes {
		if n.Name != "note1" {
			t.Fail()
		}
	}
}

func TestIsAllowed(t *testing.T) {
	if !isAllowed([]models.Group{{
		Name:   "test",
		Subnet: "127.0.0.1/32",
	}}, []string{"test"}, "127.0.0.1") {
		t.Fail()
	}

	if isAllowed([]models.Group{{
		Name:   "test",
		Subnet: "127.0.0.1/32",
	}}, []string{"test"}, "xxxxxx") {
		t.Fail()
	}

	if isAllowed([]models.Group{{
		Name:   "test",
		Subnet: "xxxxxxxxxxx",
	}}, []string{"test"}, "127.0.0.1") {
		t.Fail()
	}
}
