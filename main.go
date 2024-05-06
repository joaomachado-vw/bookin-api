package main

import (
	"log"
	"net/http"
	"time"

	"go.mod/handler"
)

const generalTimeoutSec = 15

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
	mux := http.NewServeMux()
	addRoutes(mux)

	if err := startServer(mux); err != nil {
		log.Fatal(err)
	}
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /stats", handler.BookingHandler)
}

func startServer(mux *http.ServeMux) error {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  generalTimeoutSec * time.Second,
		WriteTimeout: generalTimeoutSec * time.Second,
	}

	return s.ListenAndServe()
}
