package template

import "embed"

// TemplateFS - Contains the HTML template filesystem
//
//go:embed views/*
var TemplateFS embed.FS
