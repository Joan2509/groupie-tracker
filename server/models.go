package server

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
	Dates []string `json:"dates"`
}

type Loc struct {
	Locations []string `json:"locations"`
}

type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}
