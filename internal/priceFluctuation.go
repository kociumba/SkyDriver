package internal

import (
	"math"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
)

// Constants for weights - adjust these based on observed performance
const (
	// w1 = 0.05   // Weight for Price Spread
	// w2 = 0.2375 // Weight for Volume Imbalance
	// w3 = 0.2375 // Weight for Order Imbalance
	// w4 = 0.2375 // Weight for Moving Week Trend
	// w5 = 0.2375 // Weight for Top Order Book Pressure

	w1 = 0.18 // Weight for Price Spread
	w2 = 0.18 // Weight for Volume Imbalance
	w3 = 0.18 // Weight for Order Imbalance
	w4 = 0.18 // Weight for Moving Week Trend
	w5 = 0.18 // Weight for Top Order Book Pressure
	w6 = 0.10 // Weight for Volume Factor (new)

	// Constants for Volume Factor calculation
	lowVolumeThreshold  = 2000   // Adjust based on your data
	highVolumeThreshold = 200000 // Adjust based on your data
)

// PriceSpread calculates the percentage spread between buy and sell prices
func PriceSpread(product api.Product) float64 {
	return (product.QuickStatus.BuyPrice - product.QuickStatus.SellPrice) / product.QuickStatus.SellPrice * 100
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

// VolumeFactor calculates a factor based on weekly buy and sell volumes
func VolumeFactor(product api.Product) float64 {
	totalVolume := float64(product.QuickStatus.BuyMovingWeek + product.QuickStatus.SellMovingWeek)

	if totalVolume <= lowVolumeThreshold {
		return -100 // Maximum negative impact
	} else if totalVolume >= highVolumeThreshold {
		return 100 // Maximum positive impact
	}

	// Linear interpolation between low and high thresholds
	return -100 + (totalVolume-lowVolumeThreshold)*200/(highVolumeThreshold-lowVolumeThreshold)
}

// PredictPriceChange calculates the profit prediction
func PredictPriceChange(product api.Product) (float64, float64) {
	ps := PriceSpread(product)
	vi := VolumeImbalance(product)
	oi := OrderImbalance(product)
	mwt := MovingWeekTrend(product)
	tobp := TopOrderBookPressure(product)
	vf := VolumeFactor(product) // New factor

	prediction := w1*ps + w2*vi + w3*oi + w4*mwt + w5*tobp + w6*vf

	// Calculate confidence
	sameSignCount := 0
	factors := []float64{ps, vi, oi, mwt, tobp, vf}
	for _, factor := range factors {
		if math.Signbit(prediction) == math.Signbit(factor) {
			sameSignCount++
		}
	}
	confidence := float64(sameSignCount) / float64(len(factors)) * 100

	log.Debug("Prediction calculated",
		"prediction", prediction,
		"confidence", confidence,
		"priceSpread", ps,
		"volumeImbalance", vi,
		"orderImbalance", oi,
		"movingWeekTrend", mwt,
		"topOrderBookPressure", tobp,
		"volumeFactor", vf)

	return prediction, confidence
}
