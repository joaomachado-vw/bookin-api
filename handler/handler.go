package handler

import (
	"net/http"

	bookings "go.mod/Bookings"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	var b = bookings.BookingList{}

	w.Header().Set("Content-Type", "application/json")
	b.W = w
	b.R = r
	b.ConvertBookings()
}

func MaximizeStats(w http.ResponseWriter, r *http.Request) {
	var b = bookings.BookingList{}

	w.Header().Set("Content-Type", "application/json")
	b.W = w
	b.R = r
	b.ConvertBookings()
	b.Maximize()
}
