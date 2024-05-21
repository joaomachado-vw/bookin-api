package bookings

import (
	"encoding/json"
	"net/http"
)

type BookingList struct {
	Bookings []Booking
	W        http.ResponseWriter
	R        *http.Request
	average  StatsResponseJSON
}

func (b BookingList) ConvertBookings() {
	err := json.NewEncoder(b.W).Encode(b.Bookings)
	if err != nil {
		ErrorResponse(b.W, err, http.StatusBadRequest)
	}

	b.BookingRequest()

	if b.R.Method == http.MethodPost {
		b.GetAverage()
	}
}

func (b BookingList) BookingRequest() {
	for _, bookingRequest := range b.Bookings {
		dateConflict := bookingRequest.CheckDateConflict()

		if !dateConflict {
			b.Bookings = append(b.Bookings, bookingRequest)
		}
	}
}

func ErrorResponse(w http.ResponseWriter, err error, errorType int) {
	http.Error(w, err.Error(), errorType)
}

func (b BookingList) GetAverage() {
	for _, bookingRequest := range b.Bookings {
		b.average.AvgNight += bookingRequest.CheckAvgNight()
		b.average.MinNight = bookingRequest.GetMinNight()
		b.average.MaxNight = bookingRequest.GetMaxNight()
	}

	b.average.AvgNight /= float64(len(b.Bookings))
}

func (b BookingList) Maximize() {
	const hundred = 100

	for _, booking := range b.Bookings {
		booking.maximize.RequestIDs = append(booking.maximize.RequestIDs, booking.booking.RequestID)
		booking.maximize.TotalProfit += booking.booking.SellingRate * (booking.booking.Margin / hundred)
		booking.maximize.AvgNight = booking.booking.SellingRate *
			((booking.booking.Margin / hundred) / float64(booking.booking.Nights))
		booking.maximize.MinNight = booking.GetMinNight()
		booking.maximize.MaxNight = booking.GetMaxNight()
		booking.maximize.AvgNight /= float64(len(b.Bookings))
	}
}
