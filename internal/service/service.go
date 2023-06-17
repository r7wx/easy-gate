package service

import "html/template"

// Service - Service model
type Service struct {
	Icon     template.URL
	Name     string
	URL      template.URL
	Category string
	Groups   []string
}
