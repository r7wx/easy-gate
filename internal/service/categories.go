package service

import (
	"github.com/r7wx/easy-gate/internal/config"
)

func (s *Service) getCategories(cfg *config.Config) []category {
	categories := []category{}
	for _, cfgCategory := range cfg.Categories {
		category := category{
			Id: cfgCategory.Id,
			Title: cfgCategory.Title,
			Description:  cfgCategory.Description,
		}
		categories = append(categories, category)
	}
	return categories
}
