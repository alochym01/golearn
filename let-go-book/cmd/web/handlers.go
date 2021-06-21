package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/alochym01/golearn/pkg/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/".
	// If
	//  - it doesn't the http.NotFound() function to send a 404 response to the client.
	// else
	//  - return Hello from Snippetbox
	if r.URL.Path != "/" {
		// http.NotFound(w, r)

		// using Helper function
		app.notFound(w)

		// using return to exit home function
		return
	}

	// Use the template.ParseFiles() function to read the template file into a template set.
	// If there's an error
	// 	- We log the detailed error message
	// 	- Using http.Error() function to send a generic 500 Internal Server Error

	// Initialize a slice containing the paths to the two files.
	// The order should be
	// 	- home.page.tmpl file must be the *first* file in the slice.
	// 	- base.layout.tmpl
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// It’s important to point out that the file path
	// 	- template.ParseFiles() function must either be
	// 		- relative to your current working directory, or an absolute path.
	// We use the path relative to the root of the project directory.
	ts, err := template.ParseFiles(files...)

	if err != nil {
		// log.Println(err.Error())

		// http.Error(w, "Internal Server Error", 500)

		// using dependency inject errorLog from app
		// app.errorLog.Println(err.Error())

		// using Helper function
		app.serverError(w, err)

		// using return to exit home function
		return
	}

	// We then use the Execute() method
	// 	- To write the template content as the response body.
	// 	- The last parameter to Execute() represents dynamic data that we want to pass in,
	// which for now we'll leave as nil.
	err = ts.Execute(w, nil)

	if err != nil {
		// log.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)

		// using dependency inject errorLog from app
		// app.errorLog.Println(err.Error())

		// using Helper function
		app.serverError(w, err)
	}
}

func (app *application) showAlochym(w http.ResponseWriter, r *http.Request) {
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
		// http.NotFound(w, r)

		// using Helper function
		app.notFound(w)

		// using return to exit showAlochym function
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")

		// In contrast, the Add() method
		//  - Appends a new "Cache-Control" header.
		//  - Can be called multiple times.
		w.Header().Add("Cache-Control", "public")

		// If you want to send a non-200 status code and a plain-text response body
		// Use the http.Error() function
		//  - To send a 405 status code and "Method Not Allowed" string as the response body.
		// http.Error(w, "Method Not Allowed", 405)

		// using Helper function
		app.clientError(w, http.StatusMethodNotAllowed)

		// using return to exit showSnippet function
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// w.Write([]byte("Show Snippetbox"))

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%v", *s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
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
		// w.WriteHeader(405)

		// w.Write() will automatically send a 200 OK status code to the user
		// w.Write([]byte("Method Not Allowed"))

		// using Helper function
		app.clientError(w, http.StatusMethodNotAllowed)

		// using return to exit createSnippet function
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	// w.Write([]byte("Create Snippetbox"))
}
