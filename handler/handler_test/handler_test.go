package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	bookings "go.mod/Bookings"
	"go.mod/handler"
)

func TestHandlerGET(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/stats", nil)
	response := httptest.NewRecorder()
	handler.BookingHandler(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestHandlerPOST(t *testing.T) {
	t.Parallel()

	expectedBody := `
	  {
		"request_id": "test_1",
		"check_in": "2020-01-01",
		"nights": 5,
		"selling_rate": 200,
		"margin": 20
	  }`
	body := []byte(expectedBody)
	ctx := context.TODO()

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	handler.BookingHandler(response, req)

	var (
		actualBooking bookings.BookingList
	)

	err := json.Unmarshal(response.Body.Bytes(), &actualBooking)
	if err != nil {
		t.Fatal(err)
	}

	var expectedBooking bookings.BookingList

	err = json.Unmarshal([]byte(expectedBody), &expectedBooking)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedBooking, actualBooking)
}

//func TestStatsResponse(t *testing.T) {
//	t.Parallel()
//
//	bodyJSON := `
//	[
//	  {
//		"request_id": "test_1",
//		"check_in": "2020-01-01",
//		"nights": 5,
//		"selling_rate": 200,
//		"margin": 20
//	  },
//
//	{
//		"request_id": "test_2",
//		"check_in": "2020-01-10",
//		"nights": 5,
//		"selling_rate": 100,
//		"margin": 10
//	  }
//	]
//	`
//
//	body := []byte(bodyJSON)
//	ctx := context.TODO()
//	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/stats", bytes.NewBuffer(body))
//	response := httptest.NewRecorder()
//	handler.BookingHandler(response, req)
//	expectedAvg := Bookings.StatsResponseJSON{
//		AvgNight: 5,
//		MinNight: 2,
//		MaxNight: 8,
//	}
//
//	expectedJSON, err := json.Marshal(expectedAvg)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	assert.Equal(t, string(expectedJSON), statsJSON)
//}
//
//func TestMaximize(t *testing.T) {
//	t.Parallel()
//
//	bodyJSON := `
//	[
//	  {
//		"request_id": "test_1",
//		"check_in": "2020-01-01",
//		"nights": 5,
//		"selling_rate": 200,
//		"margin": 20
//	  },
//
//	{
//		"request_id": "test_2",
//		"check_in": "2020-01-10",
//		"nights": 5,
//		"selling_rate": 100,
//		"margin": 10
//	  }
//	]
//	`
//
//	body := []byte(bodyJSON)
//	ctx := context.TODO()
//	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/maximize", bytes.NewBuffer(body))
//	response := httptest.NewRecorder()
//	maxProfit := handler.Maximize(response, req)
//	expectedAvg := handler.MaximizeProfitJSON{
//		RequestIDs:  []string{"test_1", "test_2"},
//		TotalProfit: 50,
//		AvgNight:    5,
//		MinNight:    2,
//		MaxNight:    8,
//	}
//
//	expectedJSON, err := json.Marshal(expectedAvg)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	assert.Equal(t, string(expectedJSON), maxProfit)
//}
