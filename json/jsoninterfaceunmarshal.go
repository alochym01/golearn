// Work with JSON.
// Convert a JSON Object into an interface.
// Step 1: Declare an empty interface.
// Step 2: Declare a JSON string.
// Step 3: Convert a JSON Object into an empty interface.

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
	var m []map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}

	fmt.Printf("%T\n", m[0])  //map[string]interface {}
	fmt.Printf("%v\n", m[0])  //map[Color:RED Title:Casablanca Year:1942]
	fmt.Printf("%#v\n", m[0]) //map[string]interface {}{"Color":"RED",  "Title":"Casablanca", "Year":1942}

	fmt.Printf("%T\n", m)  //[]map[string]interface {}
	fmt.Printf("%v\n", m)  //[map[Color:RED Title:Casablanca Year:1942]]
	fmt.Printf("%#v\n", m) //[]map[string]interface {}{map[string]interface {}  {"Color":"RED", "Title":"Casablanca", "Year":1942}}

	// Access to a struct attribute by using `.` operator
	fmt.Println(m[0]["Color"]) // Access to m[0]["Color"] intefce attribute
	fmt.Println(m[0]["Title"]) // Access to m[0]["Title"] intefce attribute
	fmt.Println(m[0]["Year"])  // Access to m[0]["Year"] intefce attribute
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
	var m []map[string]interface{}
	// Convert string into io.Reader using strings.NewReader(data)
	jsonDecoder := json.NewDecoder(strings.NewReader(data))
	err := jsonDecoder.Decode(&m)
	// err := json.NewDecoder(strings.NewReader(data)).Decode(m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}

	fmt.Printf("%T\n", m[0])  //map[string]interface {}
	fmt.Printf("%v\n", m[0])  //map[Color:RED Title:Casablanca Year:1942]
	fmt.Printf("%#v\n", m[0]) //map[string]interface {}{"Color":"RED",  "Title":"Casablanca", "Year":1942}

	fmt.Printf("%T\n", m)  //[]map[string]interface {}
	fmt.Printf("%v\n", m)  //[map[Color:RED Title:Casablanca Year:1942]]
	fmt.Printf("%#v\n", m) //[]map[string]interface {}{map[string]interface {}  {"Color":"RED", "Title":"Casablanca", "Year":1942}}

	// Access to a struct attribute by using `.` operator
	fmt.Println(m[0]["Color"]) // Access to m[0]["Color"] intefce attribute
	fmt.Println(m[0]["Title"]) // Access to m[0]["Title"] intefce attribute
	fmt.Println(m[0]["Year"])  // Access to m[0]["Year"] intefce attribute
}
