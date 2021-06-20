# Golang Project Structure
```bash
├── cmd
│   └── web
│       ├── handlers.go
│       └── main.go
├── go.mod
├── go.sum
├── main.go
├── pkg
└── ui
    ├── html
    └── static

To run a program
    - go run cmd/web/*
    - go run cmd/web/* -addr=":9999" // if using flag
```

1. The cmd directory will contain the ***application-specific*** code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the cmd/web directory.
1. The pkg directory will contain the ancillary *non-application-specific* code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the SQL database models for the project.
1. The ui directory will contain the *user-interface assets* used by the web application. Specifically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images).

## So why are we using this structure?
1. It gives a clean separation between Go and non-Go assets. All the Go code we write will live exclusively under the cmd and pkg directories, leaving the project root free to hold non-Go assets like UI files, makefiles and module definitions (including our go.mod file). This can make things easier to manage when it comes to building and deploying your application in the future.
1. It scales really nicely if you want to add another executable application to your project. For example, you might want to add a CLI (Command Line Interface) to automate some administrative tasks in the future. With this structure, you could create this CLI application
under cmd/cli and it will be able to import and reuse all the code you’ve written under the pkg directory.

### Using http.Handler Interface
1. http.Handler define
   ```go
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
   ```
1. This basically means that to be a handler an object must have a ServeHTTP() method with the exact signature
1. The simplest form a handler
    ```go
    // in this case it’s a home struct.
    // but it could equally be a string or function or anything else.
    type home struct {}

    // implemented a method with the signature ServeHTTP
    func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is my home page"))
    }

    // register this with a servemux using the Handle method
    mux := http.NewServeMux()
    mux.Handle("/", &home{})
    ```
### Using Handler Functions
1. Create a normal function
    ```go
    // home doesn’t have a ServeHTTP() method.
    func home(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is my home page"))
    }

    mux := http.NewServeMux()
    // we need to transform it into a handler using the http.HandlerFunc() adapter
    // The http.HandlerFunc() adapter works by automatically adding a ServeHTTP() method to the home function
    mux.Handle("/", http.HandlerFunc(home)) // or mux.HandleFunc("/", home)
    ```