package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerGET(t *testing.T) {
	req, _ := http.NewRequest("GET", "/stats", nil)
	response := httptest.NewRecorder()
	BookingHandler(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestHandlerPOST(t *testing.T) {
	expectedBody := `
	  {
		"request_id": "test_1",
		"check_in": "2020-01-01",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  }`
	body := []byte(expectedBody)
	req, _ := http.NewRequest("POST", "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	BookingRequestHandler(response, req)
	var actualBooking BookingRequestJSON
	json.Unmarshal(response.Body.Bytes(), &actualBooking)
	var expectedBooking BookingRequestJSON
	json.Unmarshal([]byte(expectedBody), &expectedBooking)
	assert.Equal(t, expectedBooking, actualBooking)
}

func TestBookingRequestListHandler(t *testing.T) {
	expectedBody := `
	[
	{
		"request_id": "test_1",
		"check_in": "2020-01-01",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  },
	{
		"request_id": "test_2",
		"check_in": "2020-01-10",
		"nights": 5,
		"selling_rate": 100,
		"margin": 10
	  }
	]
	`
	body := []byte(expectedBody)
	req, _ := http.NewRequest("POST", "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	BookingRequestListHandler(response, req)
	var actualBookings []BookingRequestJSON
	json.Unmarshal(response.Body.Bytes(), &actualBookings)
	var expectedBookings []BookingRequestJSON
	json.Unmarshal([]byte(expectedBody), &expectedBookings)
	assert.Equal(t, expectedBookings, actualBookings)
}

func TestStatsResponse(t *testing.T) {
	bodyJSON := `
	[
	  {
		"request_id": "test_1",
		"check_in": "2020-01-01",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  },
	 
	{
		"request_id": "test_2",
		"check_in": "2020-01-10",
		"nights": 5,
		"selling_rate": 100,
		"margin": 10
	  }
	]
	`

	body := []byte(bodyJSON)
	req, _ := http.NewRequest("POST", "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	statsJSON := StatsResponse(response, req)
	expectedAvg := StatsResponseJSON{
		AvgNight: 5,
		MinNight: 2,
		MaxNight: 8,
	}
	expectedJSON, _ := json.Marshal(expectedAvg)
	assert.Equal(t, string(expectedJSON), statsJSON)
}

func TestMaximize(t *testing.T) {
	bodyJSON := `
	[
	  {
		"request_id": "test_1",
		"check_in": "2020-01-01",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  },
	 
	{
		"request_id": "test_2",
		"check_in": "2020-01-10",
		"nights": 5,
		"selling_rate": 100,
		"margin": 10
	  }
	]
	`

	body := []byte(bodyJSON)
	req, _ := http.NewRequest("POST", "/maximize", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	maxProfit := Maximize(response, req)
	expectedAvg := MaximizeProfitJSON{
		RequestIDs:  []string{"test_1", "test_2"},
		TotalProfit: 50,
		Avg_night:   5,
		Min_night:   2,
		Max_night:   8,
	}
	expectedJSON, _ := json.Marshal(expectedAvg)
	assert.Equal(t, string(expectedJSON), maxProfit)
}
