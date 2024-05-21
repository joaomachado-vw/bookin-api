package bookings

import (
	"time"
)

type MaximizeProfitJSON struct {
	RequestIDs  []string `json:"request_ids"`
	TotalProfit float64  `json:"total_profit"`
	AvgNight    float64  `json:"avg_night"`
	MinNight    float64  `json:"min_night"`
	MaxNight    float64  `json:"max_night"`
}

type StatsResponseJSON struct {
	AvgNight float64 `json:"avg_night"`
	MinNight float64 `json:"min_night"`
	MaxNight float64 `json:"max_night"`
}

type BookingRequestJSON struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

type Booking struct {
	booking  BookingRequestJSON
	dateMap  map[string]bool
	average  StatsResponseJSON
	maximize MaximizeProfitJSON
}

func (b Booking) CheckDateConflict() bool {
	checkIn, _ := time.Parse("2006-01-02", b.booking.CheckIn)
	checkOut := checkIn.AddDate(0, 0, b.booking.Nights)

	for date := checkIn; date.Before(checkOut); date = date.AddDate(0, 0, 1) {
		dateString := date.Format("2006-01-02")
		if b.dateMap[dateString] {
			return true
		}

		b.dateMap[dateString] = true
	}

	return false
}

func (b Booking) CheckAvgNight() float64 {
	const hundred = 100

	b.average.AvgNight = b.booking.SellingRate * (b.booking.Margin / hundred) / float64(b.booking.Nights)

	return b.average.AvgNight
}

func (b Booking) GetMinNight() float64 {
	if b.average.AvgNight < b.average.MinNight || b.average.MinNight == 0 {
		b.average.MinNight = b.average.AvgNight
	}

	return b.average.MinNight
}

func (b Booking) GetMaxNight() float64 {
	if b.average.AvgNight < b.average.MaxNight {
		b.average.MaxNight = b.average.AvgNight
	}

	return b.average.MaxNight
}
