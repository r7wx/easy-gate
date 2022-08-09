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
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/r7wx/easy-gate/internal/config"
)

func (r *Routine) getIconData(service config.Service) string {
	if strings.HasPrefix(service.Icon, "data:") {
		return service.Icon
	}

	u, err := url.Parse(service.Icon)
	if err == nil && u.IsAbs() {
		return r.downloadIconFromURL(service.Icon)
	}

	u, err = url.Parse(service.URL)
	if err != nil {
		return ""
	}
	return r.downloadFavicon(fmt.Sprintf("%s://%s/%s", u.Scheme,
		u.Host, "favicon.ico"))
}

func (r *Routine) downloadIconFromURL(url string) string {
	resp, err := r.Client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	mimeType := http.DetectContentType(respBytes)
	if !strings.HasPrefix(mimeType, "image/") {
		return ""
	}

	return fmt.Sprintf(
		"data:%s;base64,%s", mimeType,
		base64.StdEncoding.EncodeToString(respBytes),
	)
}

func (r *Routine) downloadFavicon(url string) string {
	resp, err := r.Client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	mimeType := http.DetectContentType(respBytes)
	if mimeType != "image/x-icon" {
		return ""
	}

	return fmt.Sprintf(
		"data:image/x-icon;base64,%s",
		base64.StdEncoding.EncodeToString(respBytes),
	)
}
