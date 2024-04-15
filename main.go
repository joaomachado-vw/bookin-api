package main

import (
	"booking-api/handler"
	"github.com/gin-gonic/gin"
)

type StatsResponse struct {
	AverageNight float64 `json:"avg_night"`
	MinNight     float64 `json:"min_night"`
	MaxNight     float64 `json:"max_night"`
}

func main() {
	router := gin.Default()
	router.GET("/stats", handler.HandlerGET)
}
