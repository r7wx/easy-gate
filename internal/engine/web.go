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
