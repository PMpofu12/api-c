package model

type CryptocurrencyList struct {
	Cryptocurrencies []Cryptocurrency `json:"cryptocurrencies"`
}

type Cryptocurrency struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
