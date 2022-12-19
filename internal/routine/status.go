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
	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/group"
	"github.com/r7wx/easy-gate/internal/note"
	"github.com/r7wx/easy-gate/internal/service"
	"github.com/r7wx/easy-gate/internal/theme"
)

// Status - Status struct
type Status struct {
	Theme       theme.Theme
	Addr        string
	Title       string
	CertFile    string
	KeyFile     string
	Groups      []group.Group
	Services    []service.Service
	Notes       []note.Note
	BehindProxy bool
	UseTLS      bool
}

func (r *Routine) updateStatus(cfg *config.Config) *Status {
	return &Status{
		Theme:       cfg.Theme,
		Addr:        cfg.Addr,
		Title:       cfg.Title,
		CertFile:    cfg.CertFile,
		KeyFile:     cfg.KeyFile,
		Groups:      cfg.Groups,
		Services:    r.getServices(cfg),
		Notes:       r.getNotes(cfg),
		BehindProxy: cfg.BehindProxy,
		UseTLS:      cfg.UseTLS,
	}
}
