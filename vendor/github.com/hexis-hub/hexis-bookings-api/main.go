package main

import (
	"github.com/gorilla/mux"
	bookingsService "github.com/hexis-hub/hexis-bookings-api/bookings-service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/bookings", bookingsService.GetBookings).Methods("GET")
	router.HandleFunc("/bookings/{id}", bookingsService.GetBooking).Methods("GET")
	router.HandleFunc("/bookings/{id}", bookingsService.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings/{id}", bookingsService.DeleteBooking).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
