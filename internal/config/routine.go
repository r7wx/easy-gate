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
	"log"
	"sync"
	"time"
)

// Routine - Config routine struct
type Routine struct {
	sync.Mutex
	Error        error
	Config       *Config
	FilePath     string
	LastChecksum string
	Interval     time.Duration
}

// NewRoutine - Create new config routine
func NewRoutine(filePath string, interval time.Duration) *Routine {
	cfg, checksum, err := LoadConfig(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Routine{
		FilePath:     filePath,
		Config:       cfg,
		Interval:     interval,
		LastChecksum: checksum,
	}
}

// GetConfiguration - Get current configuration
func (r *Routine) GetConfiguration() (*Config, error) {
	defer r.Unlock()
	r.Lock()
	return r.Config, r.Error
}

// Start - Start config routine
func (r *Routine) Start() {
	for {
		cfg, checksum, err := LoadConfig(r.FilePath)
		if err != nil {
			r.Lock()
			r.Error = err
			r.Unlock()
			continue
		}

		r.Lock()
		r.Error = nil
		if checksum != r.LastChecksum {
			r.Config = cfg
		}
		r.LastChecksum = checksum
		r.Unlock()

		time.Sleep(r.Interval)
	}
}
