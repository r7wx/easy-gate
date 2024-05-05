package static

import "embed"

// StaticFS - Contains the public WWW filesystem
//
//go:embed public/*
var StaticFS embed.FS
