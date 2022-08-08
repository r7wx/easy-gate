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
	"log"
	"net/http"

	"github.com/r7wx/easy-gate/internal/routine"
)

// Service - Easy Gate service struct
type Service struct {
	Routine *routine.Routine
}

// NewService - Create a new service
func NewService(routine *routine.Routine) *Service {
	return &Service{routine}
}

// Serve - Serve application
func (s Service) Serve() {
	status, _ := s.Routine.GetStatus()

	http.HandleFunc("/api/data", s.data)
	http.HandleFunc("/", s.webFS)

	if status.UseTLS {
		log.Println("[Easy Gate] Listening for connections on", status.Addr, "(HTTPS)")
		if err := http.ListenAndServeTLS(status.Addr, status.CertFile,
			status.KeyFile, nil); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("[Easy Gate] Listening for connections on", status.Addr)
	if err := http.ListenAndServe(status.Addr, nil); err != nil {
		log.Fatal(err)
	}
}
