package service

// Service - Service model
type Service struct {
	Icon     string   `json:"icon"`
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Category string   `json:"category"`
	Groups   []string `json:"-"`
}
