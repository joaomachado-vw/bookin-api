package handler

import (
	"encoding/json"
	"net/http"
)

type BookingRequest struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

var bookings []BookingRequest

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookings)
	}

	if r.Method == http.MethodPost {
		var bookingRequest BookingRequest
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&bookingRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bookings = append(bookings, bookingRequest)
		json.NewEncoder(w).Encode(bookingRequest)
	}
}
