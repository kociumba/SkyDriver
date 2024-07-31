package internal

import (
	"math"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
)

// Constants for weights - adjust these based on observed performance
const (
	w1 = 0.2 // Weight for Price Spread
	w2 = 0.2 // Weight for Volume Imbalance
	w3 = 0.2 // Weight for Order Imbalance
	w4 = 0.2 // Weight for Moving Week Trend
	w5 = 0.2 // Weight for Top Order Book Pressure
)

// PriceSpread calculates the percentage spread between sell and buy prices
func PriceSpread(product api.Product) float64 {
	return (product.QuickStatus.SellPrice - product.QuickStatus.BuyPrice) / product.QuickStatus.BuyPrice * 100
}

// VolumeImbalance calculates the imbalance between buy and sell volumes
func VolumeImbalance(product api.Product) float64 {
	return float64(product.QuickStatus.BuyVolume-product.QuickStatus.SellVolume) / float64(product.QuickStatus.BuyVolume+product.QuickStatus.SellVolume) * 100
}

// OrderImbalance calculates the imbalance between buy and sell orders
func OrderImbalance(product api.Product) float64 {
	return float64(product.QuickStatus.BuyOrders-product.QuickStatus.SellOrders) / float64(product.QuickStatus.BuyOrders+product.QuickStatus.SellOrders) * 100
}

// MovingWeekTrend calculates the trend based on the past week's activity
func MovingWeekTrend(product api.Product) float64 {
	return float64(product.QuickStatus.BuyMovingWeek-product.QuickStatus.SellMovingWeek) / float64(product.QuickStatus.BuyMovingWeek+product.QuickStatus.SellMovingWeek) * 100
}

// TopOrderBookPressure calculates the pressure from the visible order book
func TopOrderBookPressure(product api.Product) float64 {
	var buyPressure, sellPressure float64
	for _, buy := range product.BuySummary {
		buyPressure += float64(buy.Amount) * buy.PricePerUnit
	}
	for _, sell := range product.SellSummary {
		sellPressure += float64(sell.Amount) * sell.PricePerUnit
	}
	return (buyPressure - sellPressure) / (buyPressure + sellPressure) * 100
}

// PredictPriceChange calculates the price change prediction
func PredictPriceChange(product api.Product) (float64, float64) {
	ps := PriceSpread(product)
	vi := VolumeImbalance(product)
	oi := OrderImbalance(product)
	mwt := MovingWeekTrend(product)
	tobp := TopOrderBookPressure(product)

	prediction := w1*ps + w2*vi + w3*oi + w4*mwt + w5*tobp

	// Calculate confidence
	sameSignCount := 0
	if math.Signbit(prediction) == math.Signbit(ps) {
		sameSignCount++
	}
	if math.Signbit(prediction) == math.Signbit(vi) {
		sameSignCount++
	}
	if math.Signbit(prediction) == math.Signbit(oi) {
		sameSignCount++
	}
	if math.Signbit(prediction) == math.Signbit(mwt) {
		sameSignCount++
	}
	if math.Signbit(prediction) == math.Signbit(tobp) {
		sameSignCount++
	}
	confidence := float64(sameSignCount) / 5 * 100

	log.Debug("Prediction calculated",
		"prediction", prediction,
		"confidence", confidence,
		"priceSpread", ps,
		"volumeImbalance", vi,
		"orderImbalance", oi,
		"movingWeekTrend", mwt,
		"topOrderBookPressure", tobp)

	return prediction, confidence
}
