package model

type Cryptocurrency struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CryptocurrencyListResponse struct {
	Cryptocurrencies []Cryptocurrency `json:"cryptocurrencies"`
}
