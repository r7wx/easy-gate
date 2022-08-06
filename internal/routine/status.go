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
	"github.com/r7wx/easy-gate/internal/models"
)

// Status - Status struct
type Status struct {
	Theme       models.Theme
	Addr        string
	Title       string
	CertFile    string
	KeyFile     string
	Groups      []models.Group
	Services    []models.Service
	Notes       []models.Note
	BehindProxy bool
	UseTLS      bool
}

func toStatus(cfg *config.Config) *Status {
	return &Status{
		Theme:       models.Theme(cfg.Theme),
		Addr:        cfg.Addr,
		Title:       cfg.Title,
		CertFile:    cfg.CertFile,
		KeyFile:     cfg.KeyFile,
		Groups:      cfg.Groups,
		Services:    getServices(cfg),
		Notes:       getNotes(cfg),
		BehindProxy: cfg.BehindProxy,
		UseTLS:      cfg.UseTLS,
	}
}

func getServices(cfg *config.Config) []models.Service {
	services := []models.Service{}
	for _, cfgService := range cfg.Services {
		services = append(services, models.Service{
			Icon:   getIconData(cfgService),
			Name:   cfgService.Name,
			URL:    cfgService.URL,
			Groups: cfgService.Groups,
			Health: checkHealth(cfgService),
		})
	}
	return services
}

func getNotes(cfg *config.Config) []models.Note {
	notes := []models.Note{}
	for _, cfgNote := range cfg.Notes {
		notes = append(notes, models.Note{
			Name:   cfgNote.Name,
			Text:   cfgNote.Text,
			Groups: cfgNote.Groups,
		})
	}
	return notes
}
