package main

import (
	"sort"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
	"github.com/kociumba/SkyDriver/env"
)

var (
	product string

	best = make([]api.Product, 0, 10)
)

func main() {
	env.LoadEnv()

	// huh.NewInput().Suggestions(products).Value(&product).Run()

	products := api.GetBazaar(product)

	// log.Info("Products:", "resp", products.Products)

	for _, v := range products.Products {
		if v.QuickStatus.BuyPrice > 0 && v.QuickStatus.SellPrice > 0 {
			best = updateBest(best, v)
		}
	}

	for _, v := range best {
		log.Infof("Best Product: %s, SellPrice: %e, BuyPrice: %e, Diff: %e",
			v.ProductID,
			v.QuickStatus.SellPrice,
			v.QuickStatus.BuyPrice,
			GetDiff(v))
	}
}

func updateBest(best []api.Product, v api.Product) []api.Product {
	diff := GetDiff(v)

	if len(best) < 10 {
		best = append(best, v)
	} else {
		minIndex := 0
		minDiff := GetDiff(best[minIndex])
		for i := 1; i < len(best); i++ {
			if currentDiff := GetDiff(best[i]); currentDiff < minDiff {
				minIndex = i
				minDiff = currentDiff
			}
		}

		if diff > minDiff {
			best[minIndex] = v
		}
	}

	sort.Slice(best, func(i, j int) bool {
		return GetDiff(best[i]) > GetDiff(best[j])
	})

	return best
}

func GetDiff(p api.Product) float64 {
	return p.QuickStatus.SellPrice - p.QuickStatus.BuyPrice
}
