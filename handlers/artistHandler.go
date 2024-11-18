package handle

import (
	"net/http"
	"strconv"

	"src/api"
)

// ArtistHandler handles requests for individual artist pages
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorFunc(w, 405)
		return
	}

	// Extract the artist ID from the requested URL
	ID := r.URL.Path[8:]
	id, err := strconv.Atoi(ID)
	if err != nil {
		ErrorFunc(w, 404)
		return
	}

	// Validate the artist ID range
	if id < 1 || id > len(api.AllData) {
		ErrorFunc(w, 404)
		return
	}

	// Get the artist data from the pre-fetched AllData
	type artistData struct {
		AllData api.AllDataType
	}

	result := artistData{
		AllData: api.AllData[id-1],
	}

	// Parse and execute the artist template with the result data
	ParseTemplate(w, "artist.html", result)
}
