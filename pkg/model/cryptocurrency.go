package model

type CryptocurrencyList struct {
	Cryptocurrencies []Cryptocurrency `json:"cryptocurrencies"`
}

type Cryptocurrency struct {
	ID                       string  `json:"id"`
	Symbol                   string  `json:"symbol"`
	Name                     string  `json:"name"`
	CurrentPrice             float64 `json:"current_price"`
	MarketCap                float64 `json:"market_cap"`
	TotalVolume              float64 `json:"total_volume"`
	High24h                  float64 `json:"high_24h"`
	Low24h                   float64 `json:"low_24h"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	AllTimeHigh              float64 `json:"ath"`
	AllTimeHighDate          string  `json:"ath_date"`
}
