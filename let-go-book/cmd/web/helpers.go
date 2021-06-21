package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Define serverError helper Function
// 	- stack trace to the errorLog
// 	- sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	// using dependency inject errorLog from app
	app.errorLog.Println(trace)

	// if we want to report is the file name and line number one step back in the stack trace
	// app.errorLog.Output(2, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// Define clientError helper Function
// 	- sends a specific status code
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Define notFound Function
// 	- for consistency
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
