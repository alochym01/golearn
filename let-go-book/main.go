package main

import (
	"fmt"
	"log"
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

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux.
	// Go’s servemux supports two different types of URL patterns:
	//  - fixed paths which don’t end with a trailing slash
	//  - subtree paths which do end with a trailing slash

	// In Go’s servemux, longer URL patterns are always longest match pattern
	// Go servemux
	//  - it doesn’t support semantic URLs with variables in them
	//  - it doesn’t support regexp-based patterns
	mux := http.NewServeMux()

	// Request URL paths are automatically sanitized.
	// If the request path contains:
	//  - ".." elements
	//  - repeated slashes - "//",
	// it will automatically redirect the user to an equivalent clean URL.
	// For example,
	//  - if a user makes a request to "/foo/bar/.."   -> curl localhost:4000/snippet/..
	//  - if a user makes a request to "/foo/bar/..//" -> curl -L localhost:4000/snippet/..//
	// they will automatically be sent a 301 Permanent Redirect to "/foo/baz".

	// Mux subtree paths
	//  - Subtree path patterns are matched
	//  - you can think of subtree paths as acting a bit like they have a wildcard at the end
	//  - The start of a request URL path matches the subtree path
	//  - "/" => catch all
	//  - example: "/" or "/static/**" => a wildcard at the end

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	// Mux subtree paths
	// If a subtree path has been registered and a request is received for that
	// subtree path without a trailing slash.
	// Then the user will automatically be sent a 301 Permanent Redirect
	// to the subtree path with the slash added.
	// For example,
	//  - if you have registered the subtree path /alochym/.
	//  - any request to /foo will be redirected to /alochym/. -> curl -L localhost:4000/alochym
	mux.HandleFunc("/alo/", showAlochym)

	// Mux fixed paths
	//  - the request URL path exactly matches the fixed path
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	mux.HandleFunc("/alo", showAlochym)

	// Log the out put to console
	log.Println("Starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// We pas two parameters:
	//  -  The TCP network address to listen on (in this case ":4000)
	//  -  The servemux we just created.

	// Behind the scenes, these functions register their routes
	// with something called the DefaultServeMux
	// http.ListenAndServe(":4000", nil)
	err := http.ListenAndServe(":4000", mux)

	// If http.ListenAndServe() returns an er
	// The log.Fatal() function will also call os.Exit(1) after writing the message.
	// causing the application to immediately exit.
	log.Fatal(err)
}
