package internal

import (
	"math"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
	"github.com/kociumba/SkyDriver/config"
)

var (
	cfg *config.Config
)

func InitializeWithConfig(config *config.Config) {
	cfg = config
}

type SmoothingFunction int

const (
	NoSmoothing SmoothingFunction = iota
	SigmoidSmoothing
	TanhSmoothing
	SaturatingSmoothing
	PiecewiseSmoothing
)

var CurrentSmoothingFunction SmoothingFunction

func (s SmoothingFunction) String() string {
	switch s {
	case NoSmoothing:
		return "none"
	case SigmoidSmoothing:
		return "sigmoid"
	case TanhSmoothing:
		return "tanh"
	case SaturatingSmoothing:
		return "saturating"
	case PiecewiseSmoothing:
		return "piecewise"
	default:
		return "error"
	}
}

func SigmoidSmooth(x float64, k float64) float64 {
	return 200/(1+math.Exp(-k*x)) - 100
}

func TanhSmooth(x float64, k float64) float64 {
	return 100 * math.Tanh(k*x)
}

func SaturatingSmooth(x float64, k float64) float64 {
	return 100 * x / math.Sqrt(1+k*x*x)
	// return x / math.Sqrt(1+k*x*x)
}

func PiecewiseSmooth(x float64, n float64) float64 {
	if x > 0 {
		// return 100 * x / math.Pow(1+math.Pow(x/100, n), 1/n)
		return x / math.Pow(1+math.Pow(x/100, n), 1/n)
	}
	// return -100 * -x / math.Pow(1+math.Pow(-x/100, n), 1/n)
	return -x / math.Pow(1+math.Pow(-x/100, n), 1/n)
}

// ApplySmoothing applies the selected smoothing function
// TODO: adjust steepness based on observed data
func ApplySmoothing(x float64) float64 {
	switch CurrentSmoothingFunction {
	case SigmoidSmoothing:
		return SigmoidSmooth(x, 0.1)
	case TanhSmoothing:
		return TanhSmooth(x, 0.1)
	case SaturatingSmoothing:
		return SaturatingSmooth(x, 0.01)
	case PiecewiseSmoothing:
		return PiecewiseSmooth(x, 2)
	default:
		return x // No smoothing
	}
}

// Constants for weights - adjust these based on observed performance
const (
// w1 = 0.1428571429 // Weight for Price Spread
// w2 = 0.1428571429 // Weight for Volume Imbalance
// w3 = 0.1428571429 // Weight for Order Imbalance
// w4 = 0.1428571429 // Weight for Moving Week Trend
// w5 = 0.1428571429 // Weight for Top Order Book Pressure
// w6 = 0.1428571429 // Weight for Volume Factor (new)
// w7 = 0.1428571429 // Weight for Profit Margin Factor (new)

// w1 = 0.15 // Weight for Price Spread
// w2 = 0.15 // Weight for Volume Imbalance
// w3 = 0.15 // Weight for Order Imbalance
// w4 = 0.15 // Weight for Moving Week Trend
// w5 = 0.15 // Weight for Top Order Book Pressure
// w6 = 0.10 // Weight for Volume Factor (new)
// w7 = 0.15 // Weight for Profit Margin Factor (new)
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
	return float64(product.QuickStatus.BuyMovingWeek + product.QuickStatus.SellMovingWeek)
}

// ProfitMarginFactor calculates a factor based on the profit margin as a percentage of the sell price
func ProfitMarginFactor(product api.Product) float64 {
	profitMargin := product.QuickStatus.BuyPrice - product.QuickStatus.SellPrice
	profitMarginPercentage := profitMargin / product.QuickStatus.SellPrice

	// Linear interpolation between low and high thresholds
	return profitMarginPercentage
}

// PredictPriceChange calculates the profit prediction
func PredictPriceChange(product api.Product) (float64, float64) {
	ps := ApplySmoothing(PriceSpread(product))
	vi := ApplySmoothing(VolumeImbalance(product))
	oi := ApplySmoothing(OrderImbalance(product))
	mwt := ApplySmoothing(MovingWeekTrend(product))
	tobp := ApplySmoothing(TopOrderBookPressure(product))
	vf := ApplySmoothing(VolumeFactor(product))
	pmf := ApplySmoothing(ProfitMarginFactor(product))

	weights := cfg.Prediction.Weights
	prediction := weights.PriceSpread*ps +
		weights.VolumeImbalance*vi +
		weights.OrderImbalance*oi +
		weights.MovingWeekTrend*mwt +
		weights.OrderBookPressure*tobp +
		weights.VolumeFactor*vf +
		weights.ProfitMarginFactor*pmf

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
		"smoothingFunction", CurrentSmoothingFunction,
		"prediction", prediction,
		"confidence", confidence,
		"priceSpread", ps,
		"volumeImbalance", vi,
		"orderImbalance", oi,
		"movingWeekTrend", mwt,
		"topOrderBookPressure", tobp,
		"volumeFactor", vf,
		"profitMarginFactor", pmf,
		"weights", weights)

	return prediction, confidence
}
