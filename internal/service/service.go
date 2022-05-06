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
	"net"
	"net/http"

	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/web"
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
	cfg, _ := s.ConfigRoutine.GetConfiguration()

	http.HandleFunc("/api/data", s.data)
	http.Handle("/", http.FileServer(web.GetWebFS()))

	if cfg.UseTLS {
		log.Println("[Easy Gate] Serving API on", cfg.Addr, "(HTTPS)")
		if err := http.ListenAndServeTLS(cfg.Addr, cfg.CertFile,
			cfg.KeyFile, nil); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("[Easy Gate] Serving API on", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (s Service) data(w http.ResponseWriter, req *http.Request) {
	cfg, cfgError := s.ConfigRoutine.GetConfiguration()

	reqIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if cfg.BehindProxy {
		reqIP = req.Header.Get("X-Forwarded-For")
		if reqIP == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
	log.Println("[Easy Gate] Request from", reqIP)

	response := response{
		Title:    cfg.Title,
		Icon:     cfg.Icon,
		Motd:     cfg.Motd,
		Services: s.getServices(cfg, reqIP),
		Notes:    s.getNotes(cfg, reqIP),
		Theme:    theme(cfg.Theme),
		Error:    "",
	}
	if cfgError != nil {
		response.Error = cfgError.Error()
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("[Easy Gate] Service error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
