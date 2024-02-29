package model

type MarketHistory struct {
    Prices       [][]float64 `json:"prices"`        // Historical prices for a given cryptocurrency [timestamp, price]
    MarketCaps   [][]float64 `json:"market_caps"`   // Historical market caps [timestamp, market cap]
    TotalVolumes [][]float64 `json:"total_volumes"` // Historical total volumes [timestamp, total volume]
}