package main

import (
	"fmt"
	"math"
)

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func main() {
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}
