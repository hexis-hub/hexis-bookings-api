[![Go Report Card](https://goreportcard.com/badge/github.com/hexis-hub/hexis-bookings-api)](https://goreportcard.com/report/github.com/hexis-hub/hexis-bookings-api)

# Hexis Bookings API
API to handle our internal shared items like books, games and more.

## Note
Currently under development

## Stack
- golang
- dep
- encoding/json
- github.com/gorilla/mux
- github.com/google/uuid
- time
- log

## How to run
```javascript

// Locally
dep ensure // for dependencies if never installed
go build main.go
go run main.go

// With docker

- $ docker build -t hexis/bookings-api .
- $ docker run --rm -p 3000:3000 hexis/bookings-api
```

After that, you can open the Browser or try with the postman hiting the `http://localhost:3000` where the API will be served.