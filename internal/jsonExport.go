package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/api"
)

func ExportJson(products []api.Product, priceLimit, weeklySellLimit float64, search string, MaxDisplayedItems int) {

	results := make([]api.Results, 0, len(products))

	for _, v := range products {
		prediction, confidence := PredictPriceChange(v)

		results = append(results, api.Results{
			ProductID: v.ProductID,
			SellPrice: v.QuickStatus.SellPrice,
			BuyPrice:  v.QuickStatus.BuyPrice,
			Diff:      GetDiff(v),
			WeeklyTrafic: api.WeeklyTraffic{
				Sell: float64(v.QuickStatus.SellMovingWeek),
				Buy:  float64(v.QuickStatus.BuyMovingWeek),
			},
			Prediction: prediction,
			Confidence: confidence,
		})
	}

	exportData := api.JsonExport{
		Limit:     priceLimit,
		Sell:      weeklySellLimit,
		Search:    search,
		Max:       MaxDisplayedItems,
		Smoothing: CurrentSmoothingFunction.String(),
		Date:      time.Now(),

		Results: results,
	}

	jsonData, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		log.Error("Error marshaling to JSON:", "error", err)
		return
	}

	fmt.Println(string(jsonData))
}
