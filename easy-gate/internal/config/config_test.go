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
	"testing"
	"time"
)

const (
	testConfigFilePath = "./config.json"
)

func TestMain(m *testing.M) {
	testCfg := Config{
		Groups:   []Group{},
		Services: []Service{},
		Notes:    []Note{},
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

func TestConfig(t *testing.T) {
	routine := NewRoutine(testConfigFilePath,
		8*time.Millisecond)
	go routine.Start()

	counter := 0
	for {
		if counter == 150 {
			break
		}

		newCfg := Config{
			Groups: []Group{},
			Services: []Service{
				{
					Name: time.Now().String(),
				},
			},
			Notes: []Note{},
		}
		cfgJSON, err := json.Marshal(newCfg)
		if err != nil {
			t.Fatal(err)
		}
		err = ioutil.WriteFile(testConfigFilePath,
			cfgJSON, 0644)
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(10 * time.Millisecond)
		cfg := routine.GetConfiguration()
		if cfg.Services[0].Name != newCfg.Services[0].Name {
			t.Fatalf("Expected %v, got %v",
				cfg.Services[0].Name, newCfg.Services[0].Name)
		}

		counter++
	}
}
