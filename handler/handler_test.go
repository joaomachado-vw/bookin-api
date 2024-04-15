package handler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
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
	BookingHandler(response, req)
	assert.Equal(t, expectedBody+"\n", response.Body.String())
}
