package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadTemplates(t *testing.T) {
	// Create temporary directory for test templates
	tempDir, err := os.MkdirTemp("", "test_templates")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create "templates" subdirectory
	templatesDir := filepath.Join(tempDir, "templates")
	if err := os.Mkdir(templatesDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Create test template files
	layoutPath := filepath.Join(templatesDir, "layout.html")
	pagePath := filepath.Join(templatesDir, "page.html")

	if err := os.WriteFile(layoutPath, []byte("layout"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(pagePath, []byte("page"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Set working directory to temp dir
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Test loadTemplates
	templates, err := loadTemplates()
	if err != nil {
		t.Fatalf("loadTemplates() error = %v", err)
	}
	if len(templates) != 1 {
		t.Errorf("loadTemplates() returned %d templates, want 1", len(templates))
	}
	if _, ok := templates["page.html"]; !ok {
		t.Errorf("loadTemplates() did not load page.html")
	}
}

func TestFetchData(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"name": "Test Artist"}`)
	}))
	defer ts.Close()

	// Test fetchData
	var result map[string]string
	err := fetchData(ts.URL, &result)
	if err != nil {
		t.Fatalf("fetchData() error = %v", err)
	}
	expected := map[string]string{"name": "Test Artist"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("fetchData() = %v, want %v", result, expected)
	}
}

func TestFetchArtists(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]Artist{{Name: "Test Artist"}})
	}))
	defer ts.Close()

	// Replace the real URL with the test server URL
	oldURL := artistsURL
	artistsURL = ts.URL
	defer func() { artistsURL = oldURL }()

	// Test FetchArtists
	err := FetchArtists()
	if err != nil {
		t.Fatalf("FetchArtists() error = %v", err)
	}
	if len(artists) != 1 || artists[0].Name != "Test Artist" {
		t.Errorf("FetchArtists() did not populate artists correctly")
	}
}

func TestFetchAllLocations(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(LocationsResponse{
			Locations: []Loc{
				{
					ID:        1,
					Locations: []string{"Location1", "Location2"},
				},
			},
		})
	}))
	defer ts.Close()

	// Replace the real URL with the test server URL
	oldURL := locationsURL
	locationsURL = ts.URL
	defer func() { locationsURL = oldURL }()

	// Test FetchAllLocations
	err := FetchAllLocations()
	if err != nil {
		t.Fatalf("FetchAllLocations() error = %v", err)
	}
	if len(locations.Locations) != 1 || locations.Locations[0].Locations[0] != "Location1" {
		t.Errorf("FetchAllLocations() did not populate locations correctly")
	}
}

func TestFetchAllDates(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(DatesResponse{
			Dates: []Date{
				{
					ID:    1,
					Dates: []string{"2024-12-01", "2024-12-02"},
				},
			},
		})
	}))
	defer ts.Close()

	// Replace the real URL with the test server URL
	oldURL := datesURL
	datesURL = ts.URL
	defer func() { datesURL = oldURL }()

	// Test FetchAllDates
	err := FetchAllDates()
	if err != nil {
		t.Fatalf("FetchAllDates() error = %v", err)
	}
	if len(dates.Dates) != 1 || dates.Dates[0].Dates[0] != "2024-12-01" {
		t.Errorf("FetchAllDates() did not populate dates correctly")
	}
}

func TestFetchAllRelation(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(RelationResponse{
			Relations: []Relation{
				{
					ID: 1,
					DatesLocation: map[string][]string{
						"2024-12-01": {"Location1", "Location2"},
					},
				},
			},
		})
	}))
	defer ts.Close()

	// Replace the real URL with the test server URL
	oldURL := relationURL
	relationURL = ts.URL
	defer func() { relationURL = oldURL }()

	// Test FetchAllRelation
	err := FetchAllRelation()
	if err != nil {
		t.Fatalf("FetchAllRelation() error = %v", err)
	}
	if len(relations.Relations) != 1 || relations.Relations[0].DatesLocation["2024-12-01"][0] != "Location1" {
		t.Errorf("FetchAllRelation() did not populate relations correctly")
	}
}
