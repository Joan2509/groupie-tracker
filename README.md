# Groupie-Tracker-Visualizations

## Overview

Groupie Tracker visualizations consist of manipulating data obtained from various APIs and presenting it in an engaging and accessible manner. The application allows users to track their favorite music groups, view upcoming concerts, and their locations.

## Features
1. **User-Friendly Interface**: Intuitive navigation and layout for easy access to information.
2. **Dynamic Data Visualization**: Display information using various formats like cards, and graphics.
3. **Responsive Design**: The website is designed to be seamlessly adaptive and accessible on various devices.
4. **Error Handling**: Ensure the application handles errors gracefully and provides meaningful feedback to users.
5. **Search Functionality**: Allows users to quickly find specific artists/bands by typing on the search bar.
6. **Reduction of short-term memory load**: Minimize the amount of information users need to remember by providing clear instructions and context.
7. **Consistency**: Ensure that all elements behave similarly across the application.

## Technologies Used
- Backend: Go (Golang) for server-side logic and API interaction.
- HTML/CSS: For structuring and styling the web application.
- JavaScript: For dynamic content manipulation and API interactions.
- APIs: To fetch real-time data about music groups and concerts.

## Set up and Usage
Ensure you have Go downloaded and installed in your machine. To use the program:
1. Clone the repository:
```bash
   git clone https://github.com/Joan2509/groupie-tracker-visualizations.git
```
2. Navigate to the project directory and run the application by respectively typing the following commands:
```bash
cd groupie-tracker-visualizations
go run main.go
```
3. Open your browser and navigate to `http://localhost:3000/`

## Project Structure
The project is organized into several packages:
1. `server`: Contains files that handle:
    -   Storing the data structures representing the API data such as artist, location, date, and relation.
    -   API requests and data fetching.
    -   Managing the HTTP routes and handlers for serving the web pages and processing requests.
2. `templates`: Contains the HTML templates for the homepage, artist information page and error page.
3. `static`: Contains CSS and other static assets for the website.

## Error Handling
- `400: Bad Request`: Returned when the server cannot process the request due to a client error say empty searches.
- `404: Not Found`: Returned when the server can not find the requested resource/page.
- `405: Method Allowed`: Returned when the HTTP method used is not supported for the specified resource.
- `500: Internal Server Error`: Generic error message returned when an issue is encountered on the server side. 

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Authors
[Hillary Okello](https://learn.zone01kisumu.ke/git/hilaokello/)

[Joan Wambugu](https://learn.zone01kisumu.ke/git/jwambugu/)