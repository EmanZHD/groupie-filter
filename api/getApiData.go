package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// url api
const URL = "https://groupietrackers.herokuapp.com/api"

// ArtistType represents the data structure for an artist
type ArtistType struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// LocationType represents the data structure for an locations
type LocationType struct {
	Index []struct {
		Id       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

// DateType represents the data structure for dates
type DateType struct {
	Index []struct {
		Id   int      `json:"id"`
		Date []string `json:"dates"`
	} `json:"index"`
}
type Date struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

// RelationType represents the data structure for relations
type RelationType struct {
	Index []struct {
		Id       int                 `json:"id"`
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
type Relation struct {
	Id       int                 `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

// AllDataType combines all data types for a complete artist result
type AllDataType struct {
	Artist   ArtistType
	Location Location
	Date     Date
	Relation Relation
}

var (
	ArtistData   []ArtistType
	LocationData LocationType
	DateData     DateType
	RelationData RelationType
	AllData      []AllDataType
)

// GetApiData fetches data (json) from a given API and decodes it into the provided data structure
func GetApiData(url string, data any) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(data)
}

func init() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		GetApiData(URL+"/artists", &ArtistData)
	}()

	go func() {
		defer wg.Done()
		GetApiData(URL+"/locations", &LocationData)
	}()

	go func() {
		defer wg.Done()
		GetApiData(URL+"/dates", &DateData)
	}()

	go func() {
		defer wg.Done()
		GetApiData(URL+"/relation", &RelationData)
	}()
	wg.Wait()

	// Combine all data into FullArtistsData
	for i := 0; i < len(ArtistData); i++ {
		data := AllDataType{
			Artist:   ArtistData[i],
			Location: LocationData.Index[i],
			Date:     DateData.Index[i],
			Relation: RelationData.Index[i],
		}
		AllData = append(AllData, data)
	}
}
