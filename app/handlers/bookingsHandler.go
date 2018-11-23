package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	. "github.com/hexis-hub/hexis-bookings-api/app/models"
)

//initialize Mock Data
var bookings = []Booking{
	Booking{
		UUID: uuid.New(),
		User: &User{
			UUID:  uuid.New(),
			Name:  "Mark Ruffalo",
			Email: "hulk@smash.it",
		},
		Items: []*Item{
			&Item{
				UUID:       uuid.New(),
				Name:       "ES6 & Beyond",
				Type:       "book",
				PickedUpAt: time.Now().String(),
				ReturnAt:   time.Now().String(),
			},
		},
	},
}

// GetBookings - Retrieves all bookings
func GetBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// GetBooking - Retrieves a single booking
func GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range bookings {
		if item.UUID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Booking{})
}

// CreateBooking - Create a booking
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booking Booking
	_ = json.NewDecoder(r.Body).Decode(&booking)
	booking.UUID = uuid.New()
	booking.User.UUID = uuid.New()
	for _, item := range booking.Items {
		item.UUID = uuid.New()
	}
	bookings = append(bookings, booking)
	json.NewEncoder(w).Encode(booking)
}

// DeleteBooking - Delete a booking
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range bookings {
		if item.UUID.String() == params["id"] {
			bookings = append(bookings[:index], bookings[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bookings)
}
