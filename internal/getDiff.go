package internal

import "github.com/kociumba/SkyDriver/api"

func GetDiff(p api.Product) float64 {
	return -(p.QuickStatus.SellPrice - p.QuickStatus.BuyPrice)
}
