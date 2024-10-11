package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templates map[string]*template.Template

const (
	artistsURL = "https://groupietrackers.herokuapp.com/api/artists"
)

var artists []Artist

func init() {
	templates = make(map[string]*template.Template)
	layout := "templates/layout.html"
	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, page := range pages {
		if page != layout {
			files := []string{layout, page}
			templates[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
		}
	}
	if err := FetchArtists(); err != nil {
		log.Fatal("Could not fetch artists")
	}
}

func FetchArtists() error {
	response, err := http.Get(artistsURL)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &artists)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func FetchLocations(url string) (Loc, error) {
	var location Loc
	response, err := http.Get(url)
	if err != nil {
		return location, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	defer response.Body.Close()
	return location, nil
}

func FetchRelation(url string) (Relation, error) {
	var rel Relation
	response, err := http.Get(url)
	if err != nil {
		return rel, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return rel, err
	}
	err = json.Unmarshal(bytes, &rel)
	if err != nil {
		return rel, err
	}
	return rel, nil
}

func FetchDates(url string) (Date, error) {
	var dates Date
	response, err := http.Get(url)
	if err != nil {
		return dates, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return dates, err
	}
	err = json.Unmarshal(bytes, &dates)
	if err != nil {
		return dates, err
	}
	return dates, nil
}

// func Search(data Everything, searchTerm string) (Everything, error) {
// 	var output Everything
// 	ids := make(map[int]int)
// 	var artists []Artist
// 	for _, result := range data.Everyone {
// 		if strings.Contains(strings.ToLower(result.Name), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(result.FirstAlbum), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(strconv.Itoa(result.CreationDate)), strings.ToLower(searchTerm)) {
// 			if _, ok := ids[result.ID]; ok {
// 				continue
// 			} else {
// 				artists = append(artists, result)
// 				ids[result.ID] += 1
// 			}
// 		}
// 	}
// 	output.Everyone = artists
// 	return output, nil
// }
