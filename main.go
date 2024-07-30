package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/kociumba/SkyDriver/api"
	"github.com/kociumba/SkyDriver/env"
	"github.com/kociumba/SkyDriver/styles"
)

var (
	product string

	best = make([]api.Product, 0, 10)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E2E2")).
		// Background(lipgloss.Color("#0c0c0c")).
		Bold(true).
		Align(lipgloss.Center)

	priceLimit = flag.Float64("limit", 1e+24, "price limit")
)

func main() {
	env.LoadEnv()

	// huh.NewInput().Suggestions(products).Value(&product).Run()

	flag.Parse()

	products := api.GetBazaar(product)

	// log.Info("Products:", "resp", products.Products)

	for _, v := range products.Products {
		if v.QuickStatus.BuyPrice > 100 &&
			v.QuickStatus.SellPrice > 100 &&
			v.QuickStatus.SellMovingWeek > 100 &&
			v.QuickStatus.BuyMovingWeek > 100 &&
			v.QuickStatus.BuyPrice < *priceLimit {
			best = updateBest(best, v)
		}
	}

	t := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#D8D8D8"))).
		Headers(
			"Product/"+styles.Faint.Render(fmt.Sprintf("price limit: %.2f ", *priceLimit)),
			"SellPrice",
			"BuyPrice",
			"Difference",
			"Weekly Trafic",
		).StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return HeaderStyle
		default:
			return lipgloss.NewStyle()
		}
	})

	row := 1
	for _, v := range best {
		t.Row(
			// HACK This is fucking stupid
			styles.ProductStyle.Render(func() string {
				if row < 10 {
					return fmt.Sprintf("%v.  %v", row, v.ProductID)
				} else {
					return fmt.Sprintf("%v. %v", row, v.ProductID)
				}
			}()),
			styles.SellPriceStyle.Render(
				fmt.Sprintf("%.2f", v.QuickStatus.SellPrice),
			),
			styles.BuyPriceStyle.Render(
				fmt.Sprintf("%.2f", v.QuickStatus.BuyPrice),
			),
			styles.DiffStyle.Render(
				fmt.Sprintf("%.2f", GetDiff(v)),
			),
			styles.WeeklychangeStyle.Render(func() string {
				s, err := styles.EqualSpacingOnDividerFromInput(
					fmt.Sprintf("Sell:%v |  Buy:%v", v.QuickStatus.SellMovingWeek, v.QuickStatus.BuyMovingWeek),
					"|",
					13,
				)
				if err != nil {
					return err.Error()
				} else {
					return s
				}
			}(),
			),
		)

		row++
	}

	fmt.Println(t.String())

	// result := strings.Trim(out.String(), "\n")
	// fmt.Println(outStyle.Render(result))
}

func GetDiff(p api.Product) float64 {
	return -(p.QuickStatus.SellPrice - p.QuickStatus.BuyPrice)
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
