# Work with JSON in GO
## Marshal a struct into JSON Object in GO
- Convert a struct into JSON Object for debug
### JSON Encode in GO
- How to
    1. Declare a Movie struct
        ```go
        // Movie struct
        type Movie struct {
        	Title string
        	Year  int
            Color string
            State bool
        }
        ```
    2. Declare a variable Movie elements of a struct
        ```go
        // Define a slice of Movie
        var movies = []Movie{
        	{Title: "Casablanca", Year: 1942, Color: "RED", State: true},
        }
        ```
    3. Convert a struct into JSON Object
        ```go
        // Option 1
        // Marshal a struct with Pretty Print
	    dataIndent, err := json.MarshalIndent(m, "", "   ")
	    if err != nil {
	    	log.Fatalf("JSON marshaling failed: %s", err)
	    }
        fmt.Printf("%s\n", dataIndent)

        // Option 2
        // Marshal a struct
	    dataNoIndent, err := json.Marshal(m)
	    if err != nil {
	    	log.Fatalf("JSON marshaling failed: %s", err)
	    }
        fmt.Printf("%s\n", dataNoIndent)

        // Option 3
        // JSONNewEncoder same as json.Marshal/json.MarshalIndent
        // but using io.Writer as the OUTPUT
        // Using json.NewEncoder is faster than json.Marshal/json.MarshalIndent
        // io.Writer are:
        // - files
        // - standard output - os.Stdout
        // - ...
	    jsonEncoder := json.NewEncoder(os.Stdout)
	    err := jsonEncoder.Encode(m)
	    // err := json.NewEncoder(os.Stdout).Encode(m)

	    if err != nil {
	    	log.Fatalf("JSON marshaling failed: %s", err)
	    }
        ```
## UnMarshal a JSON Object in GO
- There 2 type of UnMarshal JSON Object in Go
### UnMarshal using Interface Object
- We access to JSON attribute by using key/value
#### JSON Decode in GO
- How to
    1. Declare an empty interface
        ```go
        // a slice of interface
        var m []map[string]interface{}
        ```
    2. Declare a JSON string
        ```go
        data := `[{"Title":"Casablanca","Year":1942,"Color":"RED","State":true}]`
        ```
    3. Convert a JSON Object into an empty interface
        ```go
        // Option 1
        // JSONDecode same as json.UnMarshal
        // but using io.Reader as the INPUT
        // Using json.NewDecoder is faster than json.UnMarshal
        // io.Reader are:
        // - files
        // - http.Reponse.Body
        // - standard input - os.Stdin
        // - strings.NewReader
        // - ...
    	jsonDecoder := json.NewDecoder(strings.NewReader(data))
    	err := jsonDecoder.Decode(&m)
    	// err := json.NewDecoder(strings.NewReader(data)).Decode(m)
    	if err != nil {
    		log.Fatalf("JSON UnMarshaling failed: %s", err)
    	}

        // Option 2
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
        ```
### UnMarshal using Struct Object
- We can add a method for a struct
- Can using a method transform a struct value attribute
- We access to a struct attribute by using `.` operator
- We can change a struct attribute's name in JSON reponse - by using `JSON tags`
#### JSON Decode in GO
- How to
    1. Declare a Movie struct and Movie method
        ```go
        // Movie struct
        type Movie struct {
        	Title string `json:"title"`
        	Year  int    `json:"year"`
        	Color string `json:"color"`
        	State bool   `json:"status"`
        }

        // Status is a Movie method
        func (m Movie) Status() string {
        	if m.State {
        		return "Enable"
        	}
        	return "Disable"
        }
        ```
    2. Declare a JSON string
        ```go
        data := `[{"Title":"Casablanca","Year":1942,"Color":"RED","State":true}]`
        ```
    3. Convert a JSON Object into a Movie struct
        ```go
        m := []Movie{}

        // Option 1
        // JSONDecode same as json.UnMarshal
        // but using io.Reader as the INPUT
        // Using json.NewDecoder is faster than json.UnMarshal
        // io.Reader are:
        // - files
        // - http.Reponse.Body
        // - standard input - os.Stdin
        // - strings.NewReader
        // - ...
       	jsonDecoder := json.NewDecoder(strings.NewReader(data))
	    err := jsonDecoder.Decode(&m)
	    // err := json.NewDecoder(strings.NewReader(data)).Decode(m)
	    if err != nil {
	    	log.Fatalf("JSON UnMarshaling failed: %s", err)
	    }

        // Option 2
	    err := json.Unmarshal([]byte(data), &m)
	    if err != nil {
	    	log.Fatalf("JSON UnMarshaling failed: %s", err)
        }

        // Using Movie method
        fmt.Println("Movie State is", movies[0].Status())

    	fmt.Printf("%T\n", m[0])  //main.Movie
    	fmt.Printf("%v\n", m[0])  //{Casablanca 1942 RED true}
    	fmt.Printf("%#v\n", m[0]) //main.Movie{Title:"Casablanca", Year:1942,   Color:"RED", State:true}

    	fmt.Printf("%T\n", m)  //[]main.Movie
    	fmt.Printf("%v\n", m)  //[{Casablanca 1942 false}]
    	fmt.Printf("%#v\n", m) //[]main.Movie{Title:"Casablanca", Year:1942, Color:"RED", State:true}

        // Access to a struct attribute by using `.` operator
    	fmt.Println(m[0].Color) // Access to Movie.Color struct attribute
    	fmt.Println(m[0].Title) // Access to Movie.Title struct attribute
        fmt.Println(m[0].Year)  // Access to Movie.Year struct attribute
        ```

