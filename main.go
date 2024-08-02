package main

import (
	"flag"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
	"github.com/kociumba/SkyDriver/internal"
	"github.com/kociumba/SkyDriver/styles"
)

const (
	MinPrice        = 100
	MinWeeklyVolume = 10
	// MaxDisplayedItems = 10
)

var (
	// best = make([]api.Product, 0, 10)

	err error

	products api.Bazaar

	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E2E2")).
		// Background(lipgloss.Color("#0c0c0c")).
		Bold(true).
		Align(lipgloss.Center)

	priceLimit        = flag.Float64("limit", 1e+32, "Set the maximum price to display")
	weeklySellLimit   = flag.Float64("sell", 1e+32, "Set the minimum weekly volume of sales to display")
	debug             = flag.Bool("dbg", false, "")
	search            = flag.String("search", "", "Search using product names")
	skip              = flag.Bool("skip", false, "Skip the prompts\n\nI know what I'm doing 😎")
	MaxDisplayedItems = flag.Int("max", 10, "Set the maximum number of items to display")
	json              = flag.Bool("json", false, "Outputs results as JSON for piping into other programs")
)

func main() {
	// env.LoadEnv()

	// huh.NewInput().Suggestions(products).Value(&product).Run()

	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	products = api.GetBazaar()

	if !*json && !*skip {
		PromptForInput()
	}

	best := filterAndSortProducts(products.Products)

	if *json {
		internal.ExportJson(best, *priceLimit, *weeklySellLimit, *search, *MaxDisplayedItems)
		return
	}
	displayResults(best)
}

func PromptForInput() {
	if *priceLimit == 1e+32 {
		var temp string
		huh.NewInput().Value(&temp).Title("Limit the maximum price").Run()
		*priceLimit, err = strconv.ParseFloat(temp, 64)
		if err != nil && *priceLimit == 0 {
			*priceLimit = 1e+32
		}
	}

	if *weeklySellLimit == 1e+32 {
		var temp string
		huh.NewInput().Value(&temp).Title("Set the minimum amount of sold items in the las 7 days").Run()
		*weeklySellLimit, err = strconv.ParseFloat(temp, 64)
		if err != nil {
			*weeklySellLimit = 1e+32
		}
	}

	if *search == "" {
		var suggestions []string
		for id := range products.Products {
			suggestions = append(suggestions, id)
		}
		huh.NewInput().Suggestions(suggestions).Title("Search for a product").Value(search).Run()
	}
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

func filterAndSortProducts(products map[string]api.Product) []api.Product {
	var filtered []api.Product
	for _, v := range products {
		if isProductEligible(v) {
			filtered = append(filtered, v)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return internal.GetDiff(filtered[i]) > internal.GetDiff(filtered[j])
	})

	if len(filtered) > *MaxDisplayedItems {
		filtered = filtered[:*MaxDisplayedItems]
	}

	return filtered
}

// The great filter v2
//
// BUG: This is more bugged than I expected, gonna have to try some stuff
func isProductEligible(p api.Product) bool {
	log.Debug("Checking if product is eligible:", "product", p.ProductID, "limit", *priceLimit, "sell", *weeklySellLimit, "search", *search)

	if *search != "" {
		if !strings.Contains(strings.ToLower(p.ProductID), strings.ToLower(*search)) {
			return false
		}
		return (p.QuickStatus.BuyPrice < *priceLimit || *priceLimit == 1e+32) &&
			(float64(p.QuickStatus.SellMovingWeek) > *weeklySellLimit || *weeklySellLimit == 1e+32)
	}

	return p.QuickStatus.BuyPrice > MinPrice &&
		p.QuickStatus.SellPrice > MinPrice &&
		p.QuickStatus.SellMovingWeek > MinWeeklyVolume &&
		p.QuickStatus.BuyMovingWeek > MinWeeklyVolume &&
		(p.QuickStatus.BuyPrice < *priceLimit || *priceLimit == 1e+32) &&
		(float64(p.QuickStatus.SellMovingWeek) > *weeklySellLimit || *weeklySellLimit == 1e+32)
}

func displayResults(best []api.Product) {
	spinner.New().Title("Crunching the data !").Run()

	t := createTable()
	for i, v := range best {
		addRowToTable(t, i+1, v, len(best))
	}
	fmt.Println(t.String())
}

func createTable() *table.Table {
	return table.New().
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
			"Prediction/Confidence",
		).StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return HeaderStyle
		default:
			return lipgloss.NewStyle()
		}
	})
}

func addRowToTable(t *table.Table, i int, v api.Product, length int) {
	t.Row(
		styles.ProductStyle.Render(func() string {
			maxLen := len(fmt.Sprintf("%d", length)) // Calculate the number of digits in the max length

			return fmt.Sprintf("%d.%*s%v", i, maxLen-len(fmt.Sprintf("%d", i))+1, "", v.ProductID)
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
				15,
			)
			if err != nil {
				return err.Error()
			} else {
				return s
			}
		}(),
		),
		func() string {
			prediction, confidence := internal.PredictPriceChange(v)
			if prediction > 0 {
				return styles.PredictionUP.Render(fmt.Sprintf("▲ %.2f/%.2f%%", prediction, confidence))
			} else if prediction < 0 {
				return styles.PredictionDown.Render(fmt.Sprintf("▼ %.2f/%.2f%%", prediction, confidence))
			} else {
				return styles.Faint.Render("N/A")
			}
		}(),
	)
}
