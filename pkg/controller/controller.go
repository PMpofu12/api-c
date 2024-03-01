package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCurrentMarketData(c *gin.Context) {
	currency := c.Param("currency")
	if currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "currency not specified"})
		return
	}
	cryptocurrencyList, err := cc.getCurrentMarketData(currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	c.JSON(http.StatusOK, cryptocurrencyList)
}

func (cc *CryptocurrencyController) GetCurrentCoinData(c *gin.Context) {
	currency := c.Param("currency")
	cryptocurrency := c.Param("cryptocurrency")
	if currency == "" || cryptocurrency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "currency or cryptocurrency not specified"})
		return
	}
	cryptocurrencyList, err := cc.getCurrentMarketData(currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	cryptocurrencyData, err := cc.getCurrentCoinData(cryptocurrency, cryptocurrencyList)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cryptocurrencyData)
}
