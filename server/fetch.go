package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	templates    map[string]*template.Template
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL  = "https://groupietrackers.herokuapp.com/api/relation"
	artists      []Artist
	locations    LocationsResponse
	dates        DatesResponse
	relations    RelationResponse
)

// init initializes templates and fetches artist data when the package is loaded.
func init() {
	var err error
	// Load HTML templates into the templates map
	templates, err = loadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	if err := FetchArtists(); err != nil {
		log.Fatal("could not fetch artists: ", err)
	}

	if err := FetchAllLocations(); err != nil {
		log.Fatal("could not fetch locations: ", err)
	}

	if err := FetchAllDates(); err != nil {
		log.Fatal("could not fetch datess: ", err)
	}

	if err := FetchAllRelation(); err != nil {
		log.Fatal("could not fetch relations: ", err)
	}
}

// loadTemplates loads HTML templates from the templates directory.
func loadTemplates() (map[string]*template.Template, error) {
	templates = make(map[string]*template.Template)
	layout := "templates/layout.html"

	// Get all HTML files in the "templates" directory
	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to load template files: %w", err)
	}

	for _, page := range pages {
		if page == layout {
			continue
		}
		// Combine layout with the current page template and parse the templates
		files := []string{layout, page}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %s: %w", page, err)
		}
		// Store the parsed template in the map using the base file name as key
		templates[filepath.Base(page)] = tmpl
	}
	return templates, nil
}

// Fetches data from the given URL and unmarshals it into the target struct.
func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}

	defer response.Body.Close()

	// Read response body into bytes slice
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s: %w", url, err)
	}

	// Unmarshal JSON into target struct
	if err := json.Unmarshal(bytes, target); err != nil {
		return fmt.Errorf("failed to unmarshal data from %s: %w", url, err)
	}

	return nil
}

// FetchArtists retrieves artist data from the defined artistsURL using fetchData
func FetchArtists() error {
	return fetchData(artistsURL, &artists)
}

// FetchAllLocations retrieves locations data from the defined locationsURL using fetchData
func FetchAllLocations() error {
	return fetchData(locationsURL, &locations)
}

// FetchAllDates retrieves dates data from the defined datesURL using fetchData
func FetchAllDates() error {
	return fetchData(datesURL, &dates)
}

// FetchAllRelation retrieves relation data from the defined relationsURL using fetchData
func FetchAllRelation() error {
	return fetchData(relationURL, &relations)
}
