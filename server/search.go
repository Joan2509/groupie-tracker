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
	if len(Query) < 2 {
		return suggestions
	}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: artist.Name,
					Name : artist.Name,
					Type:  "artist/band",
					ArtistID: artist.ID,
				},
			)
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), Query) {
				suggestions = append(suggestions,
					SearchSuggestion{
						Value: member,
						Name : artist.Name,
						Type:  "member",
						ArtistID: artist.ID,
					},
				)
			}
		}

		// Check first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: artist.FirstAlbum,
					Name : artist.Name,
					Type:  "first album",
					ArtistID: artist.ID,
				},
			)
		}
		// Check creation date
		creationDateStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDateStr, Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: creationDateStr,
					Name : artist.Name,
					Type:  "creation date",
					ArtistID: artist.ID,
				},
			)
		}
	}
	for _, artistLocation := range locations.Locations {
		for _, location := range artistLocation.Locations {
			if strings.Contains(strings.ToLower(location), Query) {
				suggestions = append(suggestions, SearchSuggestion{
					Value: location,
					Name: artists[artistLocation.ID-1].Name,
					Type:  "location",
					ArtistID: artistLocation.ID,
					
				})
			}
		}
	}

	if len(suggestions) > 10 {
		suggestions = suggestions[:10]
	}

	return suggestions
}

func PerformSearch(input string, artists []Artist) []SearchResult {
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

		// Search through actual locations
		for _, artistLocation := range locations.Locations {
			for _, location := range artistLocation.Locations {
				if strings.Contains(strings.ToLower(location), Query) {
					matchTypes = append(matchTypes, "location")
				}
			}
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

	return searchResults
}
