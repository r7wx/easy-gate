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
	"crypto/tls"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/r7wx/easy-gate/internal/config"
)

// Routine - Routine struct
type Routine struct {
	sync.Mutex
	Error        error
	Status       *Status
	Client       *http.Client
	FilePath     string
	LastChecksum string
	Interval     time.Duration
}

// NewRoutine - Create new config routine
func NewRoutine(filePath string, interval time.Duration) (*Routine, error) {
	routine := Routine{
		FilePath: filePath,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Interval: interval,
	}

	cfg, checksum, err := config.Load(filePath)
	if err != nil {
		return nil, err
	}
	routine.Status = routine.toStatus(cfg)
	routine.LastChecksum = checksum

	return &routine, nil
}

// GetStatus - Get current status
func (r *Routine) GetStatus() (*Status, error) {
	defer r.Unlock()
	r.Lock()
	return r.Status, r.Error
}

// Start - Start config routine
func (r *Routine) Start() {
	for {
		cfg, checksum, err := config.Load(r.FilePath)
		if err != nil {
			r.Lock()
			r.Error = err
			r.Unlock()
			continue
		}

		r.Lock()
		r.Error = nil
		if checksum != r.LastChecksum {
			log.Println("Detected configuration change, reloading...")
			r.Status = r.toStatus(cfg)
		}
		r.LastChecksum = checksum
		r.Unlock()

		time.Sleep(r.Interval)
	}
}
