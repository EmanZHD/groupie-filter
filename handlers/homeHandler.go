package handle

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"src/api"
)

type HomeData struct {
	Artists       []api.ArtistType
	FiltedArtists []api.AllDataType
	AllLocations  []string
	ErrorMessage  string
}
type Filter struct {
	NumMember      []string
	FirstAlbum_Min string
	FirstAlbum_Max string
	MinDate        string
	MaxDate        string
	AllLocation    []string
	Location       string
}

func FilterParams(filter *Filter) {
}

// HomeHandler handles requests for the home page and display all artists to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	filter := &Filter{}
	pageData := HomeData{}
	if r.URL.Path != "/" {
		ErrorFunc(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		ErrorFunc(w, 405)
		return
	}
	r.ParseForm()
	filter.FirstAlbum_Min = r.FormValue("firstAlbum_Min")
	filter.FirstAlbum_Max = r.FormValue("firstAlbum_Max")
	filter.NumMember = r.Form["quantity"]
	fmt.Println(filter.NumMember)
	filter.MinDate = r.FormValue("minDate")
	filter.MaxDate = r.FormValue("maxDate")
	filter.Location = r.FormValue("location")
	Location(filter)
	pageData.AllLocations = filter.AllLocation
	if len(filter.NumMember) == 0 && filter.FirstAlbum_Min == "" &&
		filter.FirstAlbum_Max == "" && filter.MinDate == "" &&
		filter.MaxDate == "" && filter.Location == "" {
		pageData.FiltedArtists = api.AllData
	} else {
		pageData = FilterData(filter)
	}
	// Parse and execute the home template with the artist data
	if len(pageData.FiltedArtists) != 0 {
		ParseTemplate(w, "home.html", pageData)
	} else {
		pageData.ErrorMessage = "Data not found"
		ParseTemplate(w, "home.html", pageData)
	}
}

func FilterData(filter *Filter) HomeData {
	pageData := HomeData{}
	var Valid bool
	NameList := []string{}
	for _, artist := range api.AllData {
		Valid = FilterByMemebers(filter, artist.Artist.Members) || len(filter.NumMember) == 0
		if Valid {
			Valid = FilterByDate(filter.FirstAlbum_Max, filter.FirstAlbum_Min, artist.Artist.FirstAlbum)
		} else {
			Valid = false
		}
		if Valid {
			Valid = FilterByDate(filter.MaxDate, filter.MinDate, strconv.Itoa(artist.Artist.CreationDate))
		} else {
			Valid = false
		}
		if Valid {
			Valid = FilterByLocation(filter, artist.Location.Location)
		} else {
			Valid = false
		}
		if Valid {
			pageData.FiltedArtists = append(pageData.FiltedArtists, artist)
			NameList = append(NameList, artist.Artist.Name)
		}
	}
	fmt.Println("filtered ", len(NameList))
	return pageData
}

func FilterByMemebers(filter *Filter, artist []string) bool {
	var Valid bool
	if len(filter.NumMember) != 0 {
		for _, memb := range filter.NumMember {
			if len(artist) == ToInt(memb) {
				Valid = true
				break
			} else {
				Valid = false
			}
		}
	} else {
		Valid = false
	}
	return Valid
}

func FilterByDate(Max string, Min string, artist1 string) bool {
	var Valid bool
	if Max == "" && Min == "" ||
		((Max != "" || Min != "") &&
			(Update(artist1) >= ToInt(Min) &&
				Update(artist1) <= ToInt(Max))) {
		Valid = true
	} else {
		Valid = false
	}
	return Valid
}

func FilterByLocation(filter *Filter, artist []string) bool {
	var Valid bool
	if strings.Contains(strings.Join(artist, " "), filter.Location) {
		Valid = true
	} else {
		Valid = false
	}
	return Valid
}

func Update(date string) int {
	newDate := (strings.Split(date, "-"))
	if len(newDate) == 3 {
		return ToInt(newDate[2])
	}
	return ToInt(date)
}

func ToInt(str string) int {
	newInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return newInt
}

func Location(filter *Filter) {
	for _, location := range api.LocationData.Index {
		filter.AllLocation = append(filter.AllLocation, location.Location...)
	}
}
