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
	expectedBody := `{"request_id":"test","check_in":"2024-04-29","nights":2,"selling_rate":50,"margin":20}`
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
		"request_id": "123",
		"check_in": "2024-04-29",
		"nights": 1,
		"selling_rate": 50,
		"margin": 20
	  },
	  {
		"request_id": "456",
		"check_in": "2024-05-01",
		"nights": 5,
		"selling_rate": 30,
		"margin": 30
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
		"request_id": "123",
		"check_in": "2024-04-29",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  },
	  {
		"request_id": "456",
		"check_in": "2024-05-01",
		"nights": 4,
		"selling_rate": 156,
		"margin": 22
	  }
	]
	`

	body := []byte(bodyJSON)
	req, _ := http.NewRequest("POST", "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	statsJSON := StatsResponse(response, req)
	expectedAvg := StatsResponseJSON{
		AvgNight: 8.29,
		MinNight: 8,
		MaxNight: 8.58,
	}
	actualJSON, _ := json.Marshal(expectedAvg)
	assert.Equal(t, actualJSON, statsJSON)
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
		"request_id": "test",
		"check_in": "2020-01-04",
		"nights": 4,
		"selling_rate": 156,
		"margin": 5
	  },
	{
		"request_id": "outro",
		"check_in": "2020-01-04",
		"nights": 4,
		"selling_rate": 150,
		"margin": 6
	  },
	{
		"request_id": "test_2",
		"check_in": "2020-01-10",
		"nights": 4,
		"selling_rate": 160,
		"margin": 30
	  }
	]
	`

	body := []byte(bodyJSON)
	req, _ := http.NewRequest("POST", "/maximize", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	maxProfit := Maximize(response, req)
	expectedAvg := MaximizeProfitJSON{
		requestIDs:  []string{"test_1", "test_2"},
		totalProfit: 88,
		avg_night:   10,
		min_night:   8,
		max_night:   12,
	}
	actualJSON, _ := json.Marshal(expectedAvg)
	assert.Equal(t, string(actualJSON), maxProfit)
}
