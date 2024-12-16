# Groupie Tracker-Search-Bar

Groupie Tracker Search is a web application designed to empower music enthusiasts by providing a user-friendly platform to explore and track their favorite bands and artists.This search functionality handles multiple search cases, supports real-time suggestions as the user types, and ensures case-insensitive searches. The search results include artist/band names, members, locations, first album release dates, and creation dates.

## Features

**1. Search Cases**

The program supports the following search cases:

- Artist/Band Name: Allows users to search for artists or bands by name.
    Members: Users can search for individual members of a band.
- Locations: Users can search for locations related to an artist or band's performances.
- First Album Release Date: Users can search by the release date of an artist's or band's first album.
- Creation Date: Users can search by the date the artist or band was formed.

**2. Case-Insensitive Search**

- The search input is case-insensitive, meaning it doesn't matter if the user types in uppercase or lowercase letters. The program normalizes the input and the data before performing the search.

**3. Typing Suggestions**

As users type in the search bar, they receive dynamic suggestions based on their input. The suggestions appear in real-time and display:

- Artist/Band Name
- Member Name
- Location
- First Album Release Date
- Creation Date

The suggestion list will also specify the type of result (e.g., "Phil Collins - artist/band" or "Freddie Mercury - member").

## Usage
To access the web page clone the repository using the link below;
```bash
git clone https://learn.zone01kisumu.ke/git/jwambugu/groupie-tracker-search-bar.git
```
Go to the project directory
```bash
cd groupie-tracker-search-bar
```
When on the working directory run the command below on your terminal to start the server;
```bash
go run .
```
Go to a web browser of your choice and get the web page by searching the url below;
```bash
localhost:3000
```

## How It Works

- Search Input: Users begin typing in the search bar, which triggers a fetch request to retrieve possible matches based on the input.
- Typing Suggestions: As the user types, suggestions are displayed based on the input. These suggestions are fetched dynamically and appear beneath the search bar.
- Selecting a Suggestion: When a suggestion is clicked or the user presses Enter, the user is redirected to a page displaying the relevant artist/band's detailed information.
- Search Results: If the input matches any artist, band, or related information, the program displays the results in a structured format. If no matches are found, an error page will inform the user that no results are available.

## Contribution Guidelines (if applicable)
- Fork the Repository: If you intend to contribute changes, fork the repository from github to create your own copy.

- Branching Strategy: Create a new branch for your specific development work to avoid conflicts with the main branch.

- Pull Requests: When ready, submit a pull request to propose merging your changes back into the main repository.

## Authors

This project was done by:

[@hilaokello](https://learn.zone01kisumu.ke/git/jwambugu/) 

[@jwambugu](https://learn.zone01kisumu.ke/git/jwambugu/)

## LICENSE

This project is licensed under the terms of the [MIT License](./LICENSE).