// Work with JSON.
// Convert a Movie Struct into a JSON Object.
// Step 1: Define a Movie struct.
// Step 2: Define variable movies.
// Step 3: JSON Encode a movie struct to a JSON Object.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Movie struct
type Movie struct {
	Title string
	Year  int
	Color string
	State bool
}

func main() {
	// Define variable movies
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: "RED", State: true},
	}
	// Marshal a movie struct into a JSON Object
	JSON(movies)
	JSONNewEncoder(movies) // faster than json.Marshal

	// Marshal a movie struct into a JSON Object with Pretty Print
	JSONIndent(movies)

}

// JSON with No Pretty Print
func JSON(m []Movie) {
	dataNoIndent, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", dataNoIndent)
	// [{"Title":"Casablanca","Year":1942,"Color":false}]

}

// JSONIndent with Pretty Print
func JSONIndent(m []Movie) {
	dataIndent, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", dataIndent)
	/*
		[
			{
				"Title": "Casablanca",
				"Year": 1942,
				"Color": false
			}
		]
	*/
}

// JSONNewEncoder same as json.Marshal/json.MarshalIndent
// but using io.Writer as the OUTPUT
// Using json.NewEncoder is faster than json.Marshal/json.MarshalIndent
// io.Writer are:
// - files
// - http.Reponse
// - standard output - os.Stdout
// - ...
func JSONNewEncoder(m []Movie) {
	jsonEncoder := json.NewEncoder(os.Stdout)
	err := jsonEncoder.Encode(m)
	// err := json.NewEncoder(os.Stdout).Encode(m)

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	// [{"Title":"Casablanca","Year":1942,"Color":false}]
}
