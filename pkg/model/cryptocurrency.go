package model

type CryptocurrencyList struct {
	Cryptocurrencies []Cryptocurrency `json:"cryptocurrencies"`
}

type Cryptocurrency struct {
	ID           string `json:"id"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	CurrentPrice int    `json:"current_price"`
	MarketCap    int    `json:"market_cap"`
	TotalVolume  int    `json:"total_volume"`
	High24h      int    `json:"24h_high"`
	Low24h       int    `json:"24_low"`
}
