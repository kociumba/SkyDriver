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

	w1 = 0.15 // Weight for Price Spread
	w2 = 0.15 // Weight for Volume Imbalance
	w3 = 0.15 // Weight for Order Imbalance
	w4 = 0.15 // Weight for Moving Week Trend
	w5 = 0.15 // Weight for Top Order Book Pressure
	w6 = 0.10 // Weight for Volume Factor (new)
	w7 = 0.15 // Weight for Profit Margin Factor (new)

	// Constants for Volume Factor calculation
	lowVolumeThreshold  = 2000    // Adjust based on data
	highVolumeThreshold = 1000000 // Adjust based on data

	// Constants for Profit Margin Factor calculation
	lowProfitMarginThreshold  = 0.001 // 0.1% profit margin
	highProfitMarginThreshold = 0.30  // 30% profit margin
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

// TopOrderBookPressure calculates the pressure from the visible orders
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

// ProfitMarginFactor calculates a factor based on the profit margin as a percentage of the sell price
func ProfitMarginFactor(product api.Product) float64 {
	profitMargin := product.QuickStatus.BuyPrice - product.QuickStatus.SellPrice
	profitMarginPercentage := profitMargin / product.QuickStatus.SellPrice

	if profitMarginPercentage <= lowProfitMarginThreshold {
		return -100 // Maximum negative impact
	} else if profitMarginPercentage >= highProfitMarginThreshold {
		return 100 // Maximum positive impact
	}

	// Linear interpolation between low and high thresholds
	return -100 + (profitMarginPercentage-lowProfitMarginThreshold)*200/(highProfitMarginThreshold-lowProfitMarginThreshold)
}

// PredictPriceChange calculates the profit prediction
func PredictPriceChange(product api.Product) (float64, float64) {
	ps := PriceSpread(product)
	vi := VolumeImbalance(product)
	oi := OrderImbalance(product)
	mwt := MovingWeekTrend(product)
	tobp := TopOrderBookPressure(product)
	vf := VolumeFactor(product)
	pmf := ProfitMarginFactor(product) // New factor

	prediction := w1*ps + w2*vi + w3*oi + w4*mwt + w5*tobp + w6*vf + w7*pmf

	// Calculate confidence
	sameSignCount := 0
	factors := []float64{ps, vi, oi, mwt, tobp, vf, pmf}
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
		"volumeFactor", vf,
		"profitMarginFactor", pmf)

	return prediction, confidence
}
