package bookings

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Booking struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	User  *User     `json:"user,omitempty"`
	Items []*Item   `json:"items,omitempty"`
}
type User struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
}
type Item struct {
	UUID       uuid.UUID `json:"uuid,omitempty"`
	Name       string    `json:"name,omitempty"`
	Type       string    `json:"type,omitempty"`
	PickedUpAt string    `json:"pickedUpAt,omitempty"`
	ReturnAt   string    `json:"returnAt,omitempty"`
}

var bookings []Booking

func GetBookings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bookings)
}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range bookings {
		if item.UUID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Booking{})
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	_ = json.NewDecoder(r.Body).Decode(&booking)
	booking.UUID, _ = uuid.NewV4()
	bookings = append(bookings, booking)
	json.NewEncoder(w).Encode(booking)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range bookings {
		if item.UUID.String() == params["id"] {
			bookings = append(bookings[:index], bookings[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(bookings)
	}
}
