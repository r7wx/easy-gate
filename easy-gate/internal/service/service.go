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
	"log"
	"net/http"

	"github.com/r7wx/easy-gate/internal/config"
)

// Service - Self Gate service struct
type Service struct {
	ConfigRoutine *config.Routine
}

// NewService - Create a new service
func NewService(cfgRoutine *config.Routine) *Service {
	return &Service{
		ConfigRoutine: cfgRoutine,
	}
}

// Serve - Serve application
func (s Service) Serve() {
	log.Println("Serving API on 0.0.0.0:8080")
	http.HandleFunc("/api/data", s.data)
	http.ListenAndServe(":8080", nil)
}

func (s Service) data(w http.ResponseWriter, req *http.Request) {
	reqAddr := req.Header.Get("X-Forwarded-For")
	if reqAddr == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	log.Println("Request from", reqAddr)
	cfg := s.ConfigRoutine.GetConfiguration()

	response := response{
		Services: s.getServices(cfg, reqAddr),
		Notes:    s.getNotes(cfg, reqAddr),
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
