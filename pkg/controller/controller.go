package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCurrentMarketData(c *gin.Context) {
	cryptocurrencyList, err := cc.getCurrentMarketData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
		return
	}
	c.JSON(http.StatusOK, cryptocurrencyList)
}

func (cc *CryptocurrencyController) GetCurrentCoinData(c *gin.Context) {
	id := c.Param("cryptocurrency")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not specified"})
		return
	}

	cryptocurrencyList, err := cc.getCurrentMarketData()
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
