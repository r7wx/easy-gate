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

	"github.com/r7wx/easy-gate/internal/models"
)

type response struct {
	Error    string           `json:"error"`
	Theme    models.Theme     `json:"theme"`
	Title    string           `json:"title"`
	Services []models.Service `json:"services"`
	Notes    []models.Note    `json:"notes"`
}

func (s Service) data(w http.ResponseWriter, req *http.Request) {
	status, cfgError := s.Routine.GetStatus()

	reqIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Println("[Easy Gate] Service error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if status.BehindProxy {
		reqIP = req.Header.Get("X-Forwarded-For")
		if reqIP == "" {
			log.Println("[Easy Gate] 400 Bad Request: X-Forwarded-For header is missing")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
	log.Printf("[Easy Gate] [%s] %s", reqIP, req.URL.Path)

	response := response{
		Title:    status.Title,
		Services: s.getServices(status, reqIP),
		Notes:    s.getNotes(status, reqIP),
		Theme:    models.Theme(status.Theme),
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
