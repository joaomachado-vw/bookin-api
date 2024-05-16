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
	AvgNight    float64  `json:"avg_night"`
	MinNight    float64  `json:"min_night"`
	MaxNight    float64  `json:"max_night"`
}

var _Bookings []BookingRequestJSON

var StatsResponse StatsResponseJSON

var MaxProfit MaximizeProfitJSON

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(_Bookings)
		if err != nil {
			ErrorResponse(w, err, http.StatusBadRequest)
		}
	}
}

func BookingRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method call", http.StatusBadRequest)
	}

	var bookingRequest BookingRequestJSON

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&bookingRequest)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
	}

	ValidationErrorResponse(w, r, bookingRequest)
	_Bookings = append(_Bookings, bookingRequest)

	if err := json.NewEncoder(w).Encode(bookingRequest); err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
	}

	_Bookings = nil
}

func BookingRequestListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method call", http.StatusBadRequest)
	}

	var bookingRequests []BookingRequestJSON

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&bookingRequests)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
	}

	for _, bookingRequest := range bookingRequests {
		ValidationErrorResponse(w, r, bookingRequest)
		dateConflict := CheckDateConflict(bookingRequest)

		if !dateConflict {
			_Bookings = append(_Bookings, bookingRequest)
		}
	}

	if err := json.NewEncoder(w).Encode(bookingRequests); err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
	}
}

// ESTA ES LA MEJOR FORMA DE HACER ESTO?!?!
func ValidationErrorResponse(w http.ResponseWriter, r *http.Request, bookingRequest BookingRequestJSON) {
	if (bookingRequest.RequestID == "" && bookingRequest.Nights == 0 && bookingRequest.Margin == 0 &&
		bookingRequest.CheckIn == "" && bookingRequest.SellingRate == 0) || r.Body == nil {
		http.Error(w, "All fields are mandatory", http.StatusBadRequest)
	}
}

func ErrorResponse(w http.ResponseWriter, err error, errorType int) {
	http.Error(w, err.Error(), errorType)
}

func StatsResponseHandler(w http.ResponseWriter, r *http.Request) string {
	const hundred = 100

	BookingRequestListHandler(w, r)

	for _, bookingRequest := range _Bookings {
		avgNight := bookingRequest.SellingRate * (bookingRequest.Margin / hundred) / float64(bookingRequest.Nights)
		StatsResponse.AvgNight += avgNight

		if avgNight < StatsResponse.MinNight || StatsResponse.MinNight == 0 {
			StatsResponse.MinNight = avgNight
		}

		if avgNight > StatsResponse.MaxNight {
			StatsResponse.MaxNight = avgNight
		}
	}

	StatsResponse.AvgNight /= float64(len(_Bookings))

	statsJSON, err := json.Marshal(StatsResponse)
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
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method call", http.StatusBadRequest)
	}

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
			_Bookings = append(_Bookings, bookingRequest)
		}
	}

	if err := json.NewEncoder(w).Encode(bookingRequests); err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
	}

	for _, booking := range _Bookings {
		const hundred = 100

		MaxProfit.RequestIDs = append(MaxProfit.RequestIDs, booking.RequestID)
		MaxProfit.TotalProfit += booking.SellingRate * (booking.Margin / hundred)
		avgNight := booking.SellingRate * (booking.Margin / hundred) / float64(booking.Nights)
		MaxProfit.AvgNight += avgNight

		if avgNight < MaxProfit.MinNight || MaxProfit.MinNight == 0 {
			MaxProfit.MinNight = avgNight
		}

		if avgNight > MaxProfit.MaxNight {
			MaxProfit.MaxNight = avgNight
		}

		MaxProfit.AvgNight /= float64(len(_Bookings))
	}

	maxProfitJSON, err := json.Marshal(MaxProfit)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)

		return ""
	}

	return string(maxProfitJSON)
}
