package main

import (
	"go.mod/handler"
	"log"
	"net/http"
)

type Booking struct {
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
	Nights      int     `json:"nights"`
}

type StatsResponse struct {
	AverageNight float64 `json:"avg_night"`
	MinNight     float64 `json:"min_night"`
	MaxNight     float64 `json:"max_night"`
}

func main() {
	http.HandleFunc("/stats", handler.BookingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
