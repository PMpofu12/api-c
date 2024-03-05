package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-market-data/pkg/model"
)

func (cc *CryptocurrencyController) getCurrentMarketData(vs_currency string) (model.CryptocurrencyList, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&order=market_cap_desc&per_page=10", vs_currency)
	response, err := http.Get(url)
	if err != nil {
		return model.CryptocurrencyList{}, err
	}
	defer response.Body.Close()
	var cryptocurrencyList model.CryptocurrencyList
	if err := json.NewDecoder(response.Body).Decode(&cryptocurrencyList.Cryptocurrencies); err != nil {
		return model.CryptocurrencyList{}, err
	}
	return cryptocurrencyList, nil
}

func (cc *CryptocurrencyController) getCurrentCoinData(id string, cryptocurrencyList model.CryptocurrencyList) (*model.Cryptocurrency, error) {
	for _, cryptocurrency := range cryptocurrencyList.Cryptocurrencies {
		if cryptocurrency.ID == id {
			return &cryptocurrency, nil
		}
	}
	return nil, fmt.Errorf("cryptocurrency with id '%s' not found", id)
}
