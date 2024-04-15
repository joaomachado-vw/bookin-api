package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGET(t *testing.T) {
	router := gin.Default()
	router.GET("/stats", HandlerGET)
	req, _ := http.NewRequest("GET", "/stats", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestHandlerPOST(t *testing.T) {
	router := gin.Default()
	router.GET("/stats", HandlerPOST)
	body := []byte(`{"request_id": "test","check_in":"2024-04-29","nights": "2","selling_rate":"50","margin":"20"}`)
	req, _ := http.NewRequest("GET", "/stats", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
