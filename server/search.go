package server

import (
	"strconv"
	"strings"
)

// GenerateSuggestions creates search suggestions based on the input
func GenerateSuggestions(input string, artists []Artist) []SearchSuggestion {
	var suggestions []SearchSuggestion

	Query := strings.ToLower(strings.TrimSpace(input))

	// If input is too short, return no suggestions
	if len(Query) < 1 {
		return suggestions
	}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), Query) {
			suggestions = append(suggestions, createSuggestion(artist.Name, "artist/band", artist))
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), Query) {
				suggestions = append(suggestions, createSuggestion(member, "member", artist))
			}
		}

		// Check first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), Query) {
			suggestions = append(suggestions, createSuggestion(artist.FirstAlbum, "first-album date", artist))
		}

		// Check creation date
		creationDateStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDateStr, Query) {
			suggestions = append(suggestions, createSuggestion(creationDateStr, "creation date", artist))
		}
	}

	//check locations
	suggestions = searchLocations(Query, suggestions)

	if len(suggestions) > 10 {
		suggestions = suggestions[:10]
	}

	return suggestions
}

// searchLocations creates search suggestions by searching through locations
func searchLocations(query string, sugg []SearchSuggestion) []SearchSuggestion {
	for _, artistLocation := range locations.Locations {
		for _, location := range artistLocation.Locations {
			if strings.Contains(strings.ToLower(location), query) {
				sugg = append(sugg, SearchSuggestion{
					Value:    location,
					Name:     artists[artistLocation.ID-1].Name,
					Type:     "location",
					ArtistID: artistLocation.ID,
				})
			}
		}
	}
	return sugg
}


// createSuggestion returns a populated SearchSuggestion
func createSuggestion(value, suggestionType string, artist Artist) SearchSuggestion {
	return SearchSuggestion{
		Value:    value,
		Name:     artist.Name,
		Type:     suggestionType,
		ArtistID: artist.ID,
	}
}

// PerformSearch Performs a search
func PerformSearch(input string) []SearchResult {
	Query := strings.ToLower(strings.TrimSpace(input))

	var searchResults []SearchResult

	for _, artist := range artists {
		matchTypes := []string{}

		// performing the specific checks
		if strings.Contains(strings.ToLower(artist.Name), Query) {
			matchTypes = append(matchTypes, "artist/band")
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), Query) {
				matchTypes = append(matchTypes, "member")
				break
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), Query) {
			matchTypes = append(matchTypes, "first album")
		}

		creationDateStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDateStr, Query) {
			matchTypes = append(matchTypes, "creation date")
		}

		// If any matches found, add to search results
		if len(matchTypes) > 0 {
			searchResults = append(searchResults, SearchResult{
				Artist:    artist,
				MatchType: matchTypes,
			})
		}
	}

	// Search through actual locations
	for _, artistLocation := range locations.Locations {
		for _, location := range artistLocation.Locations {
			if strings.Contains(strings.ToLower(location), Query) {
				searchResults = append(searchResults, SearchResult{
					Artist:    artists[artistLocation.ID-1],
					MatchType: []string{"location"},
				})
			}
		}
	}

	return searchResults
}
