package main

import (
	"flag"
	"fmt"
	"reflect"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
	"github.com/kociumba/SkyDriver/internal"
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

	priceLimit      = flag.Float64("limit", 1e+32, "price limit")
	weeklySellLimit = flag.Float64("sell", 1e+32, "price limit")
	debug           = flag.Bool("dbg", false, "debug")
)

func main() {
	// env.LoadEnv()

	// huh.NewInput().Suggestions(products).Value(&product).Run()

	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	products := api.GetBazaar(product)

	// log.Info("Products:", "resp", products.Products)

	// The great filter
	for _, v := range products.Products {
		if v.QuickStatus.BuyPrice > 100 &&
			v.QuickStatus.SellPrice > 100 &&
			v.QuickStatus.SellMovingWeek > 10 &&
			v.QuickStatus.BuyMovingWeek > 10 &&
			v.QuickStatus.BuyPrice < *priceLimit &&
			float64(v.QuickStatus.SellMovingWeek) > *weeklySellLimit {
			best = updateBest(best, v)
		}
	}

	t := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#D8D8D8"))).
		Headers(
			func() string {
				// HACK: but works great tho xd
				if DoNotRenderIfDefault(*priceLimit, 1e+32) == nil {
					return styles.ProductStyle.Render("Product")
				} else {
					return "Product/" + styles.Faint.Render(fmt.Sprintf("price limit: %.2f ", *priceLimit))
				}
			}(),
			"SellPrice",
			"BuyPrice",
			"Difference",
			"Weekly Trafic",
			"Prediction",
		).StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return HeaderStyle
		default:
			return lipgloss.NewStyle()
		}
	})

	row := 0
	for _, v := range best {
		row++
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
				fmt.Sprintf("%.2f", internal.GetDiff(v)),
			),
			styles.WeeklychangeStyle.Render(func() string {
				s, err := styles.EqualSpacingOnDividerFromInput(
					fmt.Sprintf("Sell:%v |  Buy:%v", v.QuickStatus.SellMovingWeek, v.QuickStatus.BuyMovingWeek),
					"|",
					14,
				)
				if err != nil {
					return err.Error()
				} else {
					return s
				}
			}(),
			),
			func() string {
				priceChange := internal.PredictPriceChange(v)
				if priceChange > 0 {
					return styles.PredictionUP.Render(fmt.Sprintf("▲ %.2f%%", priceChange))
				} else if priceChange < 0 {
					return styles.PredictionDown.Render(fmt.Sprintf("▼ %.2f%%", priceChange))
				} else {
					return styles.Faint.Render("N/A")
				}
			}(),
		)
	}

	fmt.Println(t.String())

	// result := strings.Trim(out.String(), "\n")
	// fmt.Println(outStyle.Render(result))
}

func updateBest(best []api.Product, v api.Product) []api.Product {
	diff := internal.GetDiff(v)

	if len(best) < 10 {
		best = append(best, v)
	} else {
		minIndex := 0
		minDiff := internal.GetDiff(best[minIndex])
		for i := 1; i < len(best); i++ {
			if currentDiff := internal.GetDiff(best[i]); currentDiff < minDiff {
				minIndex = i
				minDiff = currentDiff
			}
		}

		if diff > minDiff {
			best[minIndex] = v
		}
	}

	sort.Slice(best, func(i, j int) bool {
		return internal.GetDiff(best[i]) > internal.GetDiff(best[j])
	})

	return best
}

// NOTE: This is not needed
// I could just use == but fuck it
func DoNotRenderIfDefault(s, def interface{}) interface{} {
	if s == nil || def == nil {
		return nil
	}

	if reflect.DeepEqual(s, def) {
		return nil
	}

	return s
}
