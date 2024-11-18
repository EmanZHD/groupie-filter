package handle

import (
	"html/template"
	"net/http"
)

type errorType struct {
	StatusCode int
	Message    string
}

// ErrorFunc handles different types of HTTP errors and renders an error page
func ErrorFunc(w http.ResponseWriter, code int) {
	var errorMsg errorType
	switch code {
	case 500:
		errorMsg.StatusCode = code
		errorMsg.Message = "Internal Server Error"
	case 404:
		errorMsg.StatusCode = code
		errorMsg.Message = "Page Not Found"
	case 405:
		errorMsg.StatusCode = code
		errorMsg.Message = "Method not Allowd"
	}
	w.WriteHeader(errorMsg.StatusCode)

	// Parse the template file
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = t.Execute(w, errorMsg)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
