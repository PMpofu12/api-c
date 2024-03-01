package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCurrentMarketData(c *gin.Context) {
	currency := c.Param("currency")

	if currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency not specified"})
		return
	}

	cryptocurrencyList, err := cc.getMarketData(currency)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from CoinGecko API"})
		return
	}

	c.JSON(http.StatusOK, cryptocurrencyList)
}
