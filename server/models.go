package server

// Defines the data structuresto be fetched representing artists, locations, dates, and concert relations
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Date struct {
	ID int    `json:"id"`
	Dates []string `json:"dates"`
}

type DatesResponse struct {
    Dates []Date `json:"index"`
}

type Loc struct {
	ID int `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsResponse struct {
    Locations []Loc `json:"index"`
}

type Relation struct {
	ID        int      `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

type RelationResponse struct {
    Relations []Relation `json:"index"`
}

type SearchResult struct {
	Artist    Artist
	MatchType []string
}

type SearchSuggestion struct {
	Value string // The actual suggestion text
	Name string // The name of the artist/band
	Type  string // Type of suggestion (artist, member, location, etc.)
	ArtistID int // ID of the artist
}

// Passes dynamic data to HTML templates for rendering web pages.
type TemplateData struct {
	Title     string
	Artist    Artist
	Data      []Artist
	Locations Loc
	Dates     Date
	Concerts  Relation
	Query     string
	Results   []Artist
	Message   string
	Status    int
}
