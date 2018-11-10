package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/satori/go.uuid"
    "time"
    "log"
    "net/http"
)

type Booking struct {
    UUID      uuid.UUID `json:"uuid,omitempty"`
    User      *User  `json:"user,omitempty"`
    Items     []*Item `json:"items,omitempty"`
}
type User struct {
    UUID  uuid.UUID `json:"uuid,omitempty"`
    Name  string `json:"name,omitempty"`
    Email string `json:"email,omitempty"`
}
type Item struct {
    UUID uuid.UUID `json:"uuid,omitempty"`
    Name  string `json:"name,omitempty"`
    Type  string `json:"type,omitempty"`
    PickedUpAt  string `json:"pickedUpAt,omitempty"`
    ReturnAt  string `json:"returnAt,omitempty"`
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
    booking.UUID,_ = uuid.NewV4()
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

func main() {
    router := mux.NewRouter()
    items := []*Item{}
    // Generating unique ids for mocked data
    itemUID,_ := uuid.NewV4()
    userUID,_ := uuid.NewV4()
    bookingUID,_ := uuid.NewV4()
    
    items = append(items, &Item{
        UUID: itemUID,
        Name: "ES6 & Beyond",
        Type: "book",
        PickedUpAt: time.Now().String(),
        ReturnAt: time.Now().String(),
    })

    user := &User{ UUID: userUID, Name: "Mark Ruffalo", Email: "hulk@smash.it"}
    
    bookings = append(bookings, Booking{ UUID: bookingUID, User: user, Items: items })
    
    router.HandleFunc("/bookings", GetBookings).Methods("GET")
    router.HandleFunc("/bookings/{id}", GetBooking).Methods("GET")
    router.HandleFunc("/bookings/{id}", CreateBooking).Methods("POST")
    router.HandleFunc("/bookings/{id}", DeleteBooking).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":3000", router))
}