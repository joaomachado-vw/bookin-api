package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type BookingRequestJSON struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

type StatsResponseJSON struct {
	AvgNight float64 `json:"avg_night"`
	MinNight float64 `json:"min_night"`
	MaxNight float64 `json:"max_night"`
}

type MaximizeProfitJSON struct {
	RequestIDs  []string `json:"request_ids"`
	TotalProfit float64  `json:"total_profit"`
	Avg_night   float64  `json:"avg_night"`
	Min_night   float64  `json:"min_night"`
	Max_night   float64  `json:"max_night"`
}

var bookings []BookingRequestJSON

var statsResponse StatsResponseJSON

var maxProfit MaximizeProfitJSON

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
	bookings = nil
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
			dateConflict := CheckDateConflict(bookingRequest)
			if !dateConflict {
				bookings = append(bookings, bookingRequest)
			}
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

func StatsResponse(w http.ResponseWriter, r *http.Request) string {
	BookingRequestListHandler(w, r)
	for _, bookingRequest := range bookings {
		avg_night := bookingRequest.SellingRate * (bookingRequest.Margin / 100) / float64(bookingRequest.Nights)
		statsResponse.AvgNight += avg_night
		if avg_night < statsResponse.MinNight || statsResponse.MinNight == 0 {
			statsResponse.MinNight = avg_night
		}
		if avg_night > statsResponse.MaxNight {
			statsResponse.MaxNight = avg_night
		}
	}
	statsResponse.AvgNight /= float64(len(bookings))
	statsJSON, err := json.Marshal(statsResponse)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return ""
	}
	return string(statsJSON)
}

var dateMap = make(map[string]bool)

func CheckDateConflict(bookingRequest BookingRequestJSON) bool {
	checkIn, _ := time.Parse("2006-01-02", bookingRequest.CheckIn)
	checkOut := checkIn.AddDate(0, 0, bookingRequest.Nights)
	for date := checkIn; date.Before(checkOut); date = date.AddDate(0, 0, 1) {
		dateString := date.Format("2006-01-02")
		if dateMap[dateString] {
			return true
		}
		dateMap[dateString] = true
	}
	return false
}

func Maximize(w http.ResponseWriter, r *http.Request) string {
	if r.Method == http.MethodPost {

		var bookingRequests []BookingRequestJSON
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&bookingRequests)
		if err != nil {
			ErrorResponse(w, err, http.StatusBadRequest)
			return ""
		}
		for _, bookingRequest := range bookingRequests {
			ValidationErrorResponse(w, r, bookingRequest)
			dateConflict := CheckDateConflict(bookingRequest)
			if !dateConflict {
				bookings = append(bookings, bookingRequest)
			}
		}
		if err := json.NewEncoder(w).Encode(bookingRequests); err != nil {
			ErrorResponse(w, err, http.StatusInternalServerError)
		}
		for _, booking := range bookings {
			maxProfit.RequestIDs = append(maxProfit.RequestIDs, booking.RequestID)
			maxProfit.TotalProfit += booking.SellingRate * (booking.Margin / 100)
			avg_night := booking.SellingRate * (booking.Margin / 100) / float64(booking.Nights)
			maxProfit.Avg_night += avg_night
			if avg_night < maxProfit.Min_night || maxProfit.Min_night == 0 {
				maxProfit.Min_night = avg_night
			}
			if avg_night > maxProfit.Max_night {
				maxProfit.Max_night = avg_night
			}
		}

		maxProfit.Avg_night /= float64(len(bookings))
	}
	maxProfitJSON, err := json.Marshal(maxProfit)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
	}
	return string(maxProfitJSON)
}
