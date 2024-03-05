package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCurrentMarketData(c *gin.Context) {
	vs_currency := c.Query("vs_currency")
	if vs_currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vs_currency not specified"})
		return
	}
	cryptocurrencyList, err := cc.getCurrentMarketData(vs_currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	c.JSON(http.StatusOK, cryptocurrencyList)
}

func (cc *CryptocurrencyController) GetCurrentCoinData(c *gin.Context) {
	vs_currency := c.Query("vs_currency")
	if vs_currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vs_currency not specified"})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not specified"})
		return
	}
	cryptocurrencyList, err := cc.getCurrentMarketData(vs_currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	cryptocurrencyData, err := cc.getCurrentCoinData(id, cryptocurrencyList)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cryptocurrencyData)
}
