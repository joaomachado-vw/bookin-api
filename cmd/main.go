package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type BookingRequest struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"Nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

func main() {

	response, err := http.Get("https://app.swaggerhub.com/apis-docs/BlackfireSFL/BackendChallenge/1.0.1#/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject StatsResponse
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.AverageNight)
	fmt.Println(responseObject.MinNight)
	fmt.Println(responseObject.MaxNight)
	bookingRequest := []BookingRequest{
		{
			RequestID:   "test_id",
			CheckIn:     "2024-04-30",
			Nights:      3,
			SellingRate: 100,
			Margin:      20,
		},
	}
	fmt.Println(bookingRequest)
}
