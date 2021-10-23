// Work with JSON.
// Convert a JSON Object into a Movie Struct.
// Step 1: Declare an Movie struct.
// Step 2: Declare a JSON string.
// Step 3: Convert a JSON Object into an Movie struct.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Movie struct
type Movie struct {
	Title string
	Year  int
	Color string
	State bool
}

// Status is a Movie method
func (m Movie) Status() string {
	if m.State {
		return "Enable"
	}
	return "Disable"
}
func main() {
	// Define variable JSON Object
	data := `[{"Title":"Casablanca","Year":1942,"Color":"RED","State":true}]`
	// UnMarshal a JSON Object into a movie struct.
	JSONUnMarshal(data)

	fmt.Println("************")

	JSONDecode(data) // faster than json.UnMarshal
}

// JSONUnMarshal ...
func JSONUnMarshal(data string) {
	m := []Movie{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}

	// Using Movie method
	fmt.Println("Movie State is", m[0].Status())

	fmt.Printf("%T\n", m[0])  //main.Movie
	fmt.Printf("%v\n", m[0])  //{Casablanca 1942 false}
	fmt.Printf("%#v\n", m[0]) //main.Movie{Title:"Casablanca", Year:1942, Color:false}

	fmt.Printf("%T\n", m)  //[]main.Movie
	fmt.Printf("%v\n", m)  //[{Casablanca 1942 false}]
	fmt.Printf("%#v\n", m) //[]main.Movie{Title:"Casablanca", Year:1942, Color:false}

	fmt.Println(m[0].Color) // Access to Movie.Color struct attribute
	fmt.Println(m[0].Title) // Access to Movie.Title struct attribute
	fmt.Println(m[0].Year)  // Access to Movie.Year struct attribute
}

// JSONDecode same as json.UnMarshal
// but using io.Reader as the INPUT
// Using json.NewDecoder is faster than json.UnMarshal
// io.Reader are:
// - files
// - http.Reponse.Body
// - standard input - os.Stdin
// - strings.NewReader
// - ...
func JSONDecode(data string) {
	m := []Movie{}
	// Convert string into io.Reader using strings.NewReader(data)
	jsonDecoder := json.NewDecoder(strings.NewReader(data))
	err := jsonDecoder.Decode(&m)
	// err := json.NewDecoder(strings.NewReader(data)).Decode(m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}

	// Using Movie method
	fmt.Println("Movie State is", m[0].Status())

	fmt.Printf("%T\n", m[0])  //main.Movie
	fmt.Printf("%v\n", m[0])  //{Casablanca 1942 false}
	fmt.Printf("%#v\n", m[0]) //main.Movie{Title:"Casablanca", Year:1942, Color:false}

	fmt.Printf("%T\n", m)  //[]main.Movie
	fmt.Printf("%v\n", m)  //[{Casablanca 1942 false}]
	fmt.Printf("%#v\n", m) //[]main.Movie{Title:"Casablanca", Year:1942, Color:false}

	fmt.Println(m[0].Color) // Access to Movie.Color struct attribute
	fmt.Println(m[0].Title) // Access to Movie.Title struct attribute
	fmt.Println(m[0].Year)  // Access to Movie.Year struct attribute
}
