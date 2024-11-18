package main

import (
	"fmt"
	"net/http"

	handle "src/handlers"
)

func main() {
	// Set up pages handlers
	http.HandleFunc("/", handle.HomeHandler)
	http.HandleFunc("/artist/", handle.ArtistHandler)
	http.HandleFunc("/static/", handle.CssHandler)
	fmt.Println("the server is running on: http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
