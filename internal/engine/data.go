package engine

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/r7wx/easy-gate/internal/note"
	"github.com/r7wx/easy-gate/internal/service"
	"github.com/r7wx/easy-gate/internal/theme"
)

type response struct {
	Error      string            `json:"error"`
	Theme      theme.Theme       `json:"theme"`
	Title      string            `json:"title"`
	Categories []string          `json:"categories"`
	Services   []service.Service `json:"services"`
	Notes      []note.Note       `json:"notes"`
}

func (e Engine) data(w http.ResponseWriter, req *http.Request) {
	status, cfgError := e.Routine.GetStatus()

	reqIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Println("Engine error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if status.BehindProxy {
		reqIP = req.Header.Get("X-Forwarded-For")
		if reqIP == "" {
			log.Println("400 Bad Request: X-Forwarded-For header is missing")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
	log.Printf("[%s] %s", reqIP, req.URL.Path)

	services := getServices(status, reqIP)
	notes := getNotes(status, reqIP)
	categories := getCategories(services, notes)

	response := response{
		Title:      status.Title,
		Services:   services,
		Notes:      notes,
		Categories: categories,
		Theme:      status.Theme,
		Error:      "",
	}
	if cfgError != nil {
		response.Error = cfgError.Error()
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Engine error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
