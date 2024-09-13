package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

type Artist struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Members []string `json:"members"`
}

type Location struct {
	Location string `json:"location"`
}

type Date struct {
	Date string `json:"date"`
}

type Relation struct {
	ArtistID   int `json:"artistId"`
	LocationID int `json:"locationId"`
	DateID     int `json:"dateId"`
}

// Fetch data from the API
func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
