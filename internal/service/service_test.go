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
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/r7wx/easy-gate/internal/config"
)

const (
	testConfigFilePath = "./config.json"
)

func TestMain(m *testing.M) {
	testCfg := config.Config{
		Groups: []config.Group{
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
				Groups: []string{},
			},
			{
				Name:   "service2",
				Groups: []string{"test"},
			},
			{
				Name:   "service3",
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
		panic(err)
	}
	err = ioutil.WriteFile(testConfigFilePath,
		cfgJSON, 0644)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Remove(testConfigFilePath)
	os.Exit(exitCode)
}

func TestService(t *testing.T) {
	routine := config.NewRoutine(testConfigFilePath,
		8*time.Millisecond)
	go routine.Start()

	service := NewService(routine)
	cfg := service.ConfigRoutine.GetConfiguration()

	services := service.getServices(cfg, "192.168.1.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service2" {
			t.Fail()
		}
	}
	notes := service.getNotes(cfg, "192.168.1.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note2" {
			t.Fail()
		}
	}

	services = service.getServices(cfg, "10.1.5.1")
	for _, s := range services {
		if s.Name != "service1" && s.Name != "service3" {
			t.Fail()
		}
	}
	notes = service.getNotes(cfg, "10.1.5.1")
	for _, n := range notes {
		if n.Name != "note1" && n.Name != "note3" {
			t.Fail()
		}
	}

	services = service.getServices(cfg, "1.1.1.1")
	for _, s := range services {
		if s.Name != "service1" {
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
