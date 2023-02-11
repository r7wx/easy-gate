package engine

import (
	"github.com/r7wx/easy-gate/internal/note"
	"github.com/r7wx/easy-gate/internal/service"
)

func getCategories(services []service.Service, notes []note.Note) []string {
	categoryMap := map[string]uint{}
	for _, service := range services {
		if service.Category != "" {
			categoryMap[service.Category] = 1
		}
	}

	categories := []string{}
	for key := range categoryMap {
		categories = append(categories, key)
	}

	return categories
}
