package controller

import (
	"encoding/json"
	"net/http"

	"github.com/api-prices/pkg/model"
	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

// GetCryptocurrencies returns the list of supported cryptocurrencies from CoinGecko API
func (cc *CryptocurrencyController) GetCryptocurrencies(c *gin.Context) {
	// Make a GET request to CoinGecko API
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from CoinGecko API"})
		return
	}
	defer response.Body.Close()

	// Decode the response into a CryptocurrencyListResponse
	var cryptoResponse model.CryptocurrencyListResponse
	if err := json.NewDecoder(response.Body).Decode(&cryptoResponse.Cryptocurrencies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response from CoinGecko API"})
		return
	}

	// Return the list of cryptocurrencies
	c.JSON(http.StatusOK, cryptoResponse)
}
