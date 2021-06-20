package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/".
	// If
	//  - it doesn't the http.NotFound() function to send a 404 response to the client.
	// else
	//  - return Hello from Snippetbox
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		// using return to exit home function
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showAlochym(w http.ResponseWriter, r *http.Request) {
	// URL Query Strings -> /alochym?id=1
	// r.URL.Query().Get(id):
	// 	- Always return a string value for a parameter
	// 	- Or the empty string "" if no matching parameter exists.

	// 	- Convert it to an integer using the strconv.Atoi() function.
	// 	- If it can't be converted to an integer, or the value is less than 1.
	//  - we return a 404 not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		fmt.Println(id)
		http.NotFound(w, r)

		// using return to exit showAlochym function
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")

		// In contrast, the Add() method
		//  - Appends a new "Cache-Control" header.
		//  - Can be called multiple times.
		w.Header().Add("Cache-Control", "public")

		// If you want to send a non-200 status code and a plain-text response body
		// Use the http.Error() function
		//  - To send a 405 status code and "Method Not Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", 405)

		// using return to exit showSnippet function
		return
	}

	w.Write([]byte("Show Snippetbox"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not
	//  - Use the w.WriteHeader() method to send
	//      - a 405 status code.
	//      - Only using once per response, it can’t be changed after the status code has been written.
	//  - The w.Write() method to write a "Method Not Allowed" response body.
	//  - Using return to exit home function
	if r.Method != "POST" {

		// Use the Header().Set() method
		//  - To add an 'Allow: POST' header to the response header map.
		//  - The first parameter is the header name.
		//  - The second parameter is the header value.
		w.Header().Set("Allow", "POST")

		// if you want to send a non-200 status code,
		// you must call w.WriteHeader() before any call to w.Write().
		w.WriteHeader(405)

		// w.Write() will automatically send a 200 OK status code to the user
		w.Write([]byte("Method Not Allowed"))

		// using return to exit createSnippet function
		return
	}

	w.Write([]byte("Create Snippetbox"))
}
