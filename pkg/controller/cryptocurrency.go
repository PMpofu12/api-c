package controller

import (
	"encoding/json"
	"net/http"

	"github.com/api-prices/pkg/model"
	"github.com/gin-gonic/gin"
)

type CryptocurrencyController struct{}

func (cc *CryptocurrencyController) GetCryptocurrencies(c *gin.Context) {
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from CoinGecko API"})
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
