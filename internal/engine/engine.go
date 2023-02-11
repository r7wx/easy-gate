package engine

import (
	"log"
	"net/http"

	"github.com/r7wx/easy-gate/internal/routine"
)

// Engine - Easy Gate engine struct
type Engine struct {
	Routine *routine.Routine
}

// NewEngine - Create a new engine
func NewEngine(routine *routine.Routine) *Engine {
	return &Engine{routine}
}

// Serve - Serve application
func (e Engine) Serve() {
	status, _ := e.Routine.GetStatus()

	http.HandleFunc("/api/data", e.data)
	http.HandleFunc("/", e.webFS)

	if status.UseTLS {
		log.Println("Listening for connections on", status.Addr, "(HTTPS)")
		if err := http.ListenAndServeTLS(status.Addr, status.CertFile,
			status.KeyFile, nil); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Listening for connections on", status.Addr)
	if err := http.ListenAndServe(status.Addr, nil); err != nil {
		log.Fatal(err)
	}
}
