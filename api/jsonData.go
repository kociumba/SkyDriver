package api

import (
	"time"
)

type JsonExport struct {
	Limit     float64   `json:"limit"`
	Sell      float64   `json:"sell"`
	Search    string    `json:"search"`
	Max       int       `json:"max"`
	Smoothing string    `json:"smoothing"`
	Date      time.Time `json:"date"`

	Results []Results `json:"results"`
}

type Results struct {
	ProductID    string        `json:"product_id"`
	SellPrice    float64       `json:"sell_price"`
	BuyPrice     float64       `json:"buy_price"`
	Diff         float64       `json:"diff"`
	WeeklyTrafic WeeklyTraffic `json:"weekly_traffic"`
	Prediction   float64       `json:"prediction"`
	Confidence   float64       `json:"confidence"`
}

type WeeklyTraffic struct {
	Sell float64 `json:"sell"`
	Buy  float64 `json:"buy"`
}

// The JSON schema for this is fucked for some reason 💀
// I hate JSON, why is this shit the standard
