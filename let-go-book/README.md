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
```

1. The cmd directory will contain the ***application-specific*** code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the cmd/web directory.
1. The pkg directory will contain the ancillary *non-application-specific* code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the SQL database models for the project.
1. The ui directory will contain the *user-interface assets* used by the web application. Specifically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images).

## So why are we using this structure?
1. It gives a clean separation between Go and non-Go assets. All the Go code we write will live exclusively under the cmd and pkg directories, leaving the project root free to hold non-Go assets like UI files, makefiles and module definitions (including our go.mod file). This can make things easier to manage when it comes to building and deploying your application in the future.
1. It scales really nicely if you want to add another executable application to your project. For example, you might want to add a CLI (Command Line Interface) to automate some administrative tasks in the future. With this structure, you could create this CLI application
under cmd/cli and it will be able to import and reuse all the code you’ve written under the pkg directory.