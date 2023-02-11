package engine

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
	"github.com/r7wx/easy-gate/internal/group"
	"github.com/r7wx/easy-gate/internal/routine"
	"github.com/r7wx/easy-gate/internal/theme"
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
		Theme: theme.Theme{
			Background: "#ffffff",
			Foreground: "#000000",
		},
		Groups: []group.Group{
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
				Name:     "service1",
				URL:      "http://example.com/service1",
				Category: "One",
				Groups:   []string{},
			},
			{
				Name:     "service2",
				URL:      "http://example.com/service2",
				Category: "One",
				Groups:   []string{"test"},
			},
			{
				Name:   "service3",
				URL:    "http://example.com/service3",
				Groups: []string{"test2"},
			},
		},
		Notes: []config.Note{
			{
				Name:     "note1",
				Groups:   []string{},
				Category: "Two",
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

func TestEngine(t *testing.T) {
	routine, err := routine.NewRoutine(testConfigFilePath, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	go routine.Start()
	service := NewEngine(routine)

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

	service := NewEngine(routine)
	status, _ := service.Routine.GetStatus()

	services := getServices(status, "192.168.1.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service2" {
			t.Fail()
		}
	}

	services = getServices(status, "10.1.5.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service3" {
			t.Fail()
		}
	}

	services = getServices(status, "1.1.1.1")
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

	service := NewEngine(routine)
	status, _ := service.Routine.GetStatus()

	notes := getNotes(status, "192.168.1.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note2" {
			t.Fail()
		}
	}

	notes = getNotes(status, "10.1.5.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note3" {
			t.Fail()
		}
	}

	notes = getNotes(status, "1.1.1.1")
	for _, n := range notes {
		if n.Name != "note1" {
			t.Fail()
		}
	}
}
