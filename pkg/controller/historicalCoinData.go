package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoricalCoinDataController struct{}

func (hcd *HistoricalCoinDataController) GetHistoricalCoinData(c *gin.Context) {
	vs_currency := c.Query("vs_currency")
	dateFrom := c.GetFloat64(c.Query("from"))
	dateTo := c.GetFloat64(c.Query("to"))
	if vs_currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vs_currency not specified"})
		return
	}
	if dateFrom > dateTo {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dateFrom cannot be greater than dateTo"})
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not specified"})
		return
	}
	historicalCoinData, err := hcd.getHistoricalCoinData(vs_currency, dateFrom, dateTo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	c.JSON(http.StatusOK, historicalCoinData)
}
