# Learning Web with GoLang
## Folder Structure
```bash
├── domain
│   └── customer.go
├── go.mod
├── go.sum
├── handler
│   └── customer.go
├── router
│   └── router.go
├── service
│   └── customer.go
├── storage
│   └── adaptermockcustomer.go
└── web
    └── api
        └── main.go
```

1. `domain/customer.go`
   1. Define business Object - [sample code](web/domain/customer.go)
   2. Define `storage repository` interface - [sample code](web/domain/customer.go)
2. `storage/adaptermockcustomer.go`
   1. Define place to READ/WRITE/UPDATE/DELETE a Customer Object
   2. Implemetation all `storage repository` method - [sample code](web/storage/adaptermockcustomer.go)
3. `service/customer.go`:
   1. Define `business logic` interface - [sample code](web/service/customer.go)
   2. Implementation all `business logic method` - [sample code](web/service/customer.go)
4. `handler/customer.go` - Handler all User Request for [Customer business Object](web/handler/customer.go)
5.  `router/router.go` - Define all the route of web api in GO and [sample code](web/router/router.go)
6.  `web/api/main` - entry point of web api in GO and [sample code](web/web/api/main.go)
## Web Architechture - Hexagonal Domain