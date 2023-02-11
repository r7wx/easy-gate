package engine

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/r7wx/easy-gate/web"
)

func (s Engine) webFS(w http.ResponseWriter, req *http.Request) {
	webFS := web.GetWebFS()
	if _, err := webFS.Open(strings.TrimLeft(req.URL.Path, "/")); err != nil {
		req.URL.Path = "/"
	}

	reqIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Println("WebFS error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	status, _ := s.Routine.GetStatus()
	if status.BehindProxy {
		reqIP = req.Header.Get("X-Forwarded-For")
		if reqIP == "" {
			log.Println("400 Bad Request: X-Forwarded-For header is missing")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}

	log.Printf("[%s] %s", reqIP, req.URL.Path)
	http.FileServer(webFS).ServeHTTP(w, req)
}
