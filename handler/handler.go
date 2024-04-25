package handler

import (
	"encoding/json"
	"net/http"
)

type BookingRequestJSON struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

var bookings []BookingRequestJSON

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookings)
	}
}

func BookingRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var bookingRequest BookingRequestJSON
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&bookingRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ValidationErrorResponse(w, r, bookingRequest)
		bookings = append(bookings, bookingRequest)
		if err := json.NewEncoder(w).Encode(bookingRequest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func BookingRequestListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var bookingRequests []BookingRequestJSON
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&bookingRequests)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, bookingRequest := range bookingRequests {
			ValidationErrorResponse(w, r, bookingRequest)
			bookings = append(bookings, bookingRequest)
		}
		if err := json.NewEncoder(w).Encode(bookingRequests); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// ESTA ES LA MEJOR FORMA DE HACER ESTO?!?!
func ValidationErrorResponse(w http.ResponseWriter, r *http.Request, bookingRequest BookingRequestJSON) {
	if (bookingRequest.RequestID == "" && bookingRequest.Nights == 0 && bookingRequest.Margin == 0 &&
		bookingRequest.CheckIn == "" && bookingRequest.SellingRate == 0) || r.Body == nil {
		http.Error(w, "All fields are mandatory", http.StatusBadRequest)
		return
	}
}
