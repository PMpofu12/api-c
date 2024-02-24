package controller

import (
	"encoding/json"
	"net/http"

	"github.com/api-prices/pkg/model"
	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCryptocurrencies(c *gin.Context) {
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from CoinGecko API"})
		return
	}

	defer response.Body.Close()
	var cryptoResponse model.CryptocurrencyList

	if err := json.NewDecoder(response.Body).Decode(&cryptoResponse.Cryptocurrencies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response from CoinGecko API"})
		return
	}

	c.JSON(http.StatusOK, cryptoResponse)
}
