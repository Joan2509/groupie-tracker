package main

import (
    "log"
    "net/http"
	"groupie-tracker/backend"
)

func main() {
    http.HandleFunc("/", backend.homeHandler) // Home page
    http.HandleFunc("/artists", backend.artistsHandler)
    http.HandleFunc("/locations", backend.locationsHandler)
    http.HandleFunc("/dates", backend.datesHandler)
    http.HandleFunc("/relations", backend.relationsHandler)

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
