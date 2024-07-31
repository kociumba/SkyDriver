package api

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

type Bazaar struct {
	Success     bool               `json:"success"`
	LastUpdated int64              `json:"lastUpdated"`
	Products    map[string]Product `json:"products"`
}

type Product struct {
	ProductID   string      `json:"product_id"`
	SellSummary []SellInfo  `json:"sell_summary"`
	BuySummary  []BuyInfo   `json:"buy_summary"`
	QuickStatus QuickStatus `json:"quick_status"`
}

type SellInfo struct {
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Orders       int     `json:"orders"`
}

type BuyInfo struct {
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Orders       int     `json:"orders"`
}

type QuickStatus struct {
	ProductID      string  `json:"productId"`
	SellPrice      float64 `json:"sellPrice"`
	SellVolume     int     `json:"sellVolume"`
	SellMovingWeek int     `json:"sellMovingWeek"`
	SellOrders     int     `json:"sellOrders"`
	BuyPrice       float64 `json:"buyPrice"`
	BuyVolume      int     `json:"buyVolume"`
	BuyMovingWeek  int     `json:"buyMovingWeek"`
	BuyOrders      int     `json:"buyOrders"`
}

func GetBazaar(product string) Bazaar {
	// Construct the URL with the uuid and the environment variable KEY
	url := "https://api.hypixel.net/v2/skyblock/bazaar"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Unmarshal the JSON response
	var bazaar Bazaar
	if err := json.NewDecoder(resp.Body).Decode(&bazaar); err != nil {
		log.Fatal("Error decoding JSON:", "err", err, "resp", resp)
	}

	return bazaar
}
