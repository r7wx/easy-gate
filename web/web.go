package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed build/*
var webFS embed.FS

// GetWebFS - Get embedded frontend file system
func GetWebFS() http.FileSystem {
	fs, err := fs.Sub(webFS, "build")
	if err != nil {
		log.Fatal("Error loading embedded filesystem:", err)
	}
	return http.FS(fs)
}
