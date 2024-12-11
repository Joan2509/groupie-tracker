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
					Type:  "artist/band",
				},
			)
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), Query) {
				suggestions = append(suggestions,
					SearchSuggestion{
						Value: member,
						Type:  "member",
					},
				)
			}
		}

		// Check first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: artist.FirstAlbum,
					Type:  "first album",
				},
			)
		}

		if strings.Contains(strings.ToLower(artist.Locations), Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: artist.Locations,
					Type:  "location",
				},
			)
		}

		// Check creation date
		creationDateStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDateStr, Query) {
			suggestions = append(suggestions,
				SearchSuggestion{
					Value: creationDateStr,
					Type:  "creation date",
				},
			)
		}
	}

	// if len(suggestions) > 10 {
	// 	suggestions = suggestions[:10]
	// }

	return suggestions
}
