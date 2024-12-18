package server

import (
	"reflect"
	"testing"
)

func TestGenerateSuggestions(t *testing.T) {
	// Mock artist data for testing
	artists = []Artist{
		{
			ID:         3,
			Name:       "Pink Floyd",
			Members:    []string{"Syd Barrett", "Roger Waters", "David Gilmour"},
			FirstAlbum: "05-08-1967",
		},
		{
			ID:         8,
			Name:       "Kendrick Lamar",
			Members:    []string{"Kendrick Lamar Duckworth"},
			FirstAlbum: "22-10-2012",
		},
	}
	// Test cases
	tests := []struct {
		name     string
		input    string
		expected []SearchSuggestion
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: nil,
		},
		{
			name:  "Search by artist name",
			input: "pink",
			expected: []SearchSuggestion{
				{
					Value:    "Pink Floyd",
					Name:     "Pink Floyd",
					Type:     "artist/band",
					ArtistID: 3,
				},
			},
		},
		{
			name:  "Search by member name",
			input: "Kendrick Lamar",
			expected: []SearchSuggestion{
				{
					Value:    "Kendrick Lamar",
					Name:     "Kendrick Lamar",
					Type:     "artist/band",
					ArtistID: 8,
				},
				{
					Value:    "Kendrick Lamar Duckworth",
					Name:     "Kendrick Lamar",
					Type:     "member",
					ArtistID: 8,
				},
			},
		},
		{
			name:  "Search by first-album date",
			input: "05-08-1967",
			expected: []SearchSuggestion{
				{
					Value:    "05-08-1967",
					Name:     "Pink Floyd",
					Type:     "first-album date",
					ArtistID: 3,
				},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSuggestions(tt.input, artists)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GenerateSuggestions() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSearchLocations(t *testing.T) {
	// Test setup
	locations.Locations = []Loc{
		{
			ID:        1,
			Locations: []string{"London, UK", "Cambridge, UK"},
		},
		{
			ID:        2,
			Locations: []string{"Liverpool, UK", "London, UK"},
		},
	}

	// Test cases
	tests := []struct {
		name     string
		query    string
		initial  []SearchSuggestion
		expected []SearchSuggestion
	}{
		{
			name:    "Search for location",
			query:   "london",
			initial: nil, // Changed from empty slice to nil
			expected: []SearchSuggestion{
				{
					Value:    "London, UK",
					Name:     "Pink Floyd",
					Type:     "location",
					ArtistID: 1,
				},
				{
					Value:    "London, UK",
					Name:     "Kendrick Lamar",
					Type:     "location",
					ArtistID: 2,
				},
			},
		},
		{
			name:  "Search with existing suggestions",
			query: "liverpool",
			initial: []SearchSuggestion{{
				Value:    "SOJA",
				Name:     "SOJA",
				Type:     "artist/band",
				ArtistID: 2,
			}},
			expected: []SearchSuggestion{
				{
					Value:    "SOJA",
					Name:     "SOJA",
					Type:     "artist/band",
					ArtistID: 2,
				},
				{
					Value:    "Liverpool, UK",
					Name:     "Kendrick Lamar",
					Type:     "location",
					ArtistID: 2,
				},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchLocations(tt.query, tt.initial)
			if len(got) == 0 && len(tt.expected) == 0 {
				// Both slices are empty, test passes
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("searchLocations() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCreateSuggestion(t *testing.T) {
	artist := Artist{
		ID:           1,
		Name:         "Pink Floyd",
		Members:      []string{"David Gilmour", "Roger Waters", "Nick Mason"},
		CreationDate: 1965,
		FirstAlbum:   "The Piper at the Gates of Dawn",
	}

	tests := []struct {
		name           string
		value          string
		suggestionType string
		artist         Artist
		expected       SearchSuggestion
	}{
		{
			name:           "Create artist suggestion",
			value:          "Pink Floyd",
			suggestionType: "artist/band",
			artist:         artist,
			expected: SearchSuggestion{
				Value:    "Pink Floyd",
				Name:     "Pink Floyd",
				Type:     "artist/band",
				ArtistID: 1,
			},
		},
		{
			name:           "Create member suggestion",
			value:          "David Gilmour",
			suggestionType: "member",
			artist:         artist,
			expected: SearchSuggestion{
				Value:    "David Gilmour",
				Name:     "Pink Floyd",
				Type:     "member",
				ArtistID: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createSuggestion(tt.value, tt.suggestionType, tt.artist)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("createSuggestion() = %v, want %v", got, tt.expected)
			}
		})
	}
}
