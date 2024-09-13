package backend

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Printf("Error: not found")
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Artists handler
func artistsHandler(w http.ResponseWriter, r *http.Request) {
	var artists []Artist

	if err := fetchData(artistsURL, &artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/artists.html"))
	if err := tmpl.Execute(w, artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Locations handler
func locationsHandler(w http.ResponseWriter, r *http.Request) {
	var locations []Location

	if err := fetchData(locationsURL, &locations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/locations.html"))
	if err := tmpl.Execute(w, locations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Dates handler
func datesHandler(w http.ResponseWriter, r *http.Request) {
	var dates []Date

	if err := fetchData(datesURL, &dates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dates.html"))
	if err := tmpl.Execute(w, dates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Relations handler
func relationsHandler(w http.ResponseWriter, r *http.Request) {
	var relations []Relation

	if err := fetchData(relationsURL, &relations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/relations.html"))
	if err := tmpl.Execute(w, relations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events every 5 seconds
	for {
		// Create a message to send
		message := fmt.Sprintf("New event at %s", time.Now().Format(time.RFC3339))
		fmt.Fprintf(w, "data: %s\n\n", message)
		// Flush the response to the client
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(5 * time.Second) // Wait for 5 seconds before sending the next event
	}
}
