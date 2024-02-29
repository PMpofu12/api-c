package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-market-data/pkg/model"
	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCryptocurrencies(c *gin.Context) {

	currency := c.Param("currency")

	if currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency not specified"})
		return
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&order=market_cap_desc&per_page=10", currency)
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from CoinGecko API"})
		return
	}

	defer response.Body.Close()
	var cryptocurrencyList model.CryptocurrencyList

	if err := json.NewDecoder(response.Body).Decode(&cryptocurrencyList.Cryptocurrencies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response from CoinGecko API"})
		return
	}

	c.JSON(http.StatusOK, cryptocurrencyList)
}
