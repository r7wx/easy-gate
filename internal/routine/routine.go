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
	Error        error
	Status       *Status
	Client       *http.Client
	FilePath     string
	LastChecksum string
	Interval     time.Duration
	sync.Mutex
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
	routine.Status = routine.updateStatus(cfg)
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
			r.Status = r.updateStatus(cfg)
		}
		r.LastChecksum = checksum
		r.Unlock()

		time.Sleep(r.Interval)
	}
}
