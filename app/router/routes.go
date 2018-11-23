package router

import handler "github.com/hexis-hub/hexis-bookings-api/app/handlers"

var routes = Routes{
	Route{"GetBooks", "/bookings", "GET", handler.GetBookings},
	Route{"GetBook", "/bookings/{id}", "GET", handler.GetBooking},
	Route{"CreateBook", "/bookings", "POST", handler.CreateBooking},
	Route{"DeleteBook", "/bookings/{id}", "DELETE", handler.DeleteBooking},
}
