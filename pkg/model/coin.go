package model

type CurrentMarketData struct {
	CurrentMarketData []CurrentCoinData `json:"currentMarketData"`
}

type CurrentCoinData struct {
	ID                           string  `json:"id"`
	Symbol                       string  `json:"symbol"`
	Name                         string  `json:"name"`
	CurrentPrice                 float64 `json:"current_price"`
	MarketCap                    float64 `json:"market_cap"`
	MarketCapChangePercentage24h float64 `json:"market_cap_change_percentage_24h"`
	FullyDilutedValuation        float64 `json:"fully_diluted_valuation"`
	CirculatingSupply            float64 `json:"circulating_supply"`
	TotalSupply                  float64 `json:"total_supply"`
	MaxSupply                    float64 `json:"max_supply"`
	TotalVolume                  float64 `json:"total_volume"`
	High24h                      float64 `json:"high_24h"`
	Low24h                       float64 `json:"low_24h"`
	PriceChangePercentage24h     float64 `json:"price_change_percentage_24h"`
	AllTimeHigh                  float64 `json:"ath"`
	AllTimeHighDate              string  `json:"ath_date"`
	LastUpdated                  string  `json:"last_updated"`
}
