package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//						FLAG
	// Define a new command-line flag with the name 'addr'.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// flag.Parse() function to parse the command-line
	// 	- In the command-line flag value and assigns it to the addr variable.
	// 	- You need to call this *before* you use the addr variable
	// 	- Otherwise it will always contain the default value of ":4000".
	// 	- If any err encountered the application will be terminated.
	flag.Parse()

	//						LOG
	// Use log.New() to create a logger for writing information messages. This
	// three parameters:
	// 	- the destination to write the logs to (os.Stdout).
	// 	- the start prefix for message (INFO followed by a tab).
	// 	- additional information to include (local date and local time)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way:
	// 	- use the log.Lshortfile flag to include file name and line number to the err log.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// infoLog with file example
	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	// errorLog with file example
	// f, err = os.OpenFile("/tmp/error.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime)

	//						Go SERVEMUX
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

	// We use the http.FileServer() to serve static files or using Nginx web server
	// Create a file server which serves files out of the "./ui/static"
	// The path given to the http.Dir function is relative to root directory.
	// func (mux *ServeMux) Handle(pattern string, handler Handler)
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// All URL paths that start with "/static/" which are sent to fileserver
	// For matching paths we remove "/static" prefix before the request to the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Mux fixed paths
	//  - the request URL path exactly matches the fixed path
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	mux.HandleFunc("/alo", showAlochym)

	// Initialize a new http.Server struct.
	// We set
	// 	- The Addr
	// 	- Handler fields
	// 	- The ErrorLog field
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog, // use custom errorlog
		Handler:  mux,
	}

	// Log the out put to console
	infoLog.Printf("Starting server on %s", *addr)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pas two parameters:
	//  -  The TCP network address to listen on (in this case ":4000)
	//  -  The servemux we just created.

	// Behind the scenes, these functions register their routes
	// with something called the DefaultServeMux
	// http.ListenAndServe(":4000", nil)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()

	// If http.ListenAndServe() returns an er
	// The log.Fatal() function will also call os.Exit(1) after writing the message.
	// causing the application to immediately exit.
	errorLog.Fatal(err)
}
