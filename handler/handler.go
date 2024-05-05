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

type StatsResponseJSON struct {
	avg_night float64 `json:"avg_night"`
	min_night float64 `json:"min_night"`
	max_night float64 `json:"max_night"`
}

var bookings []BookingRequestJSON

var statsResponse StatsResponseJSON

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
			ErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		ValidationErrorResponse(w, r, bookingRequest)
		bookings = append(bookings, bookingRequest)
		if err := json.NewEncoder(w).Encode(bookingRequest); err != nil {
			ErrorResponse(w, err, http.StatusInternalServerError)
		}
	}
}

func BookingRequestListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var bookingRequests []BookingRequestJSON
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&bookingRequests)
		if err != nil {
			ErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		for _, bookingRequest := range bookingRequests {
			ValidationErrorResponse(w, r, bookingRequest)
			bookings = append(bookings, bookingRequest)
		}
		if err := json.NewEncoder(w).Encode(bookingRequests); err != nil {
			ErrorResponse(w, err, http.StatusInternalServerError)
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

func ErrorResponse(w http.ResponseWriter, err error, errorType int) {
	http.Error(w, err.Error(), errorType)
}

func StatsResponse(w http.ResponseWriter, r *http.Request) []byte {
	BookingRequestListHandler(w, r)
	for _, bookingRequest := range bookings {
		avg_night := (bookingRequest.SellingRate * (bookingRequest.Margin / 100) / float64(bookingRequest.Nights))
		statsResponse.avg_night += avg_night
		if avg_night < statsResponse.min_night || statsResponse.min_night == 0 {
			statsResponse.min_night = avg_night
		}
		if avg_night > statsResponse.max_night {
			statsResponse.max_night = avg_night
		}
	}
	statsResponse.avg_night /= float64(len(bookings))
	statsJSON, err := json.Marshal(statsResponse)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return nil
	}
	return statsJSON
}
