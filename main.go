package main

import (
	"log"
	"net/http"

	"github.com/hexis-hub/hexis-bookings-api/app/router"
)

func main() {
	router := router.New()

	log.Fatal(http.ListenAndServe(":3000", router))
}
