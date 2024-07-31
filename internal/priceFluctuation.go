package internal

import (
	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
)

// Improve this logic couse fuck me economics math is boring
func GetPriceFluctuation(product api.Product) (float64, float64) {
	// Calculate approximately how much an item is gaining or losing value in percent
	// from the difference in quick status buy and sell and the weekly turnover
	//
	// Logic being if the people buy an item more than they sell it the value will go up and vice versa

	// Calculate average buying price
	var buySum float64
	for _, buy := range product.BuySummary {
		buySum += buy.PricePerUnit
	}
	avgBuyPrice := buySum / float64(len(product.BuySummary))
	log.Debug("Average Buy Price calculated", "avgBuyPrice", avgBuyPrice)

	// Calculate price change based on current data
	var priceChange float64
	if avgBuyPrice > 0 {
		priceChange = (product.QuickStatus.SellPrice - avgBuyPrice) / avgBuyPrice * 100
		log.Debug("Price Change calculated", "priceChange", priceChange)
	} else {
		log.Debug("Average Buy Price is zero, skipping Price Change calculation")
	}

	// Calculate market trend
	var marketTrend float64
	if product.QuickStatus.BuyVolume > 0 {
		marketTrend = (float64(product.QuickStatus.SellVolume) - float64(product.QuickStatus.BuyVolume)) / float64(product.QuickStatus.BuyVolume) * 100
		log.Debug("Market Trend calculated", "marketTrend", marketTrend)
	} else {
		log.Debug("Buy Volume is zero, skipping Market Trend calculation")
	}

	return priceChange, marketTrend
}

// PredictPriceChange wraps GetPriceFluctuation and returns a single prediction percentage.
func PredictPriceChange(product api.Product) float64 {
	priceChange, marketTrend := GetPriceFluctuation(product)

	// Calculate the prediction as the weighted average of price change and market trend
	// Give more weight to market trend as it reflects demand and supply
	prediction := (priceChange + marketTrend) / 2
	log.Debug("Prediction calculated", "prediction", prediction)

	return prediction
}
