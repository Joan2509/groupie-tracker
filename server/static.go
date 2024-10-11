package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the /static/ prefix from the URL path
	filePath := path.Join("static", strings.TrimPrefix(r.URL.Path, "/static/"))

	// Check if the file exists and is not a directory
	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() {
		ErrorPage(w, http.StatusNotFound)
		return
	}

	// Check the file extension
	ext := filepath.Ext(filePath)
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		// For other file types, you might want to return a 403 Forbidden status
		ErrorPage(w, http.StatusForbidden)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
