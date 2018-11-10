package bookings

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

// Booking structure
type Booking struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	User  *User     `json:"user,omitempty"`
	Items []*Item   `json:"items,omitempty"`
}

// User structure
type User struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
}

// Item structure
type Item struct {
	UUID       uuid.UUID `json:"uuid,omitempty"`
	Name       string    `json:"name,omitempty"`
	Type       string    `json:"type,omitempty"`
	PickedUpAt string    `json:"pickedUpAt,omitempty"`
	ReturnAt   string    `json:"returnAt,omitempty"`
}

var items = []*Item{}
var bookings []Booking
var itemUID = uuid.NewV4()
var userUID = uuid.NewV4()
var bookingUID = uuid.NewV4()
// GetBookings - Retrieves all bookings
func GetBookings(w http.ResponseWriter, r *http.Request) {
	// Generating unique ids for mocked data
	items := append(items, &Item{
		UUID:       itemUID,
		Name:       "ES6 & Beyond",
		Type:       "book",
		PickedUpAt: time.Now().String(),
		ReturnAt:   time.Now().String(),
	})

	user := &User{UUID: userUID, Name: "Mark Ruffalo", Email: "hulk@smash.it"}

	bookings := append(bookings, Booking{UUID: bookingUID, User: user, Items: items})

	json.NewEncoder(w).Encode(bookings)

	return
}

// GetBooking - Retrieves a single booking
func GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range bookings {
		if item.UUID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Booking{})
	return
}

// CreateBooking - Create a booking
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	_ = json.NewDecoder(r.Body).Decode(&booking)
	booking.UUID = uuid.NewV4()
	bookings = append(bookings, booking)
	json.NewEncoder(w).Encode(booking)
	return
}

// DeleteBooking - Delete a booking
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range bookings {
		if item.UUID.String() == params["id"] {
			bookings = append(bookings[:index], bookings[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(bookings)
	}
	return
}
