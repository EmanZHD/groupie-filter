package handle

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

// ParseTemplate parses and executes an HTML template with the provided data
func ParseTemplate(w http.ResponseWriter, path string, data any) {
	// Parse the template file
	t, err := template.ParseFiles("templates/" + path)
	if err != nil {
		ErrorFunc(w, 500)
		return
	}

	// Execute the template with the provided data

	err = t.Execute(w, data)
	if err != nil {
		ErrorFunc(w, 500)
		return
	}
}

// Serve static files from the "static" directory
func CssHandler(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("." + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		ErrorFunc(w, 404)
		return
	}
	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}
