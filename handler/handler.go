package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingRequest struct {
	RequestID   string  `json:"request_id"`
	CheckIn     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

func HandlerGET(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"msg": "testGET"})
}

func HandlerPOST(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"msg": "testPOST"})
}
