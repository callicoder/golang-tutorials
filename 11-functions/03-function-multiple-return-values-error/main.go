package main

import (
	"errors"
	"fmt"
	"math"
)

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func main() {
	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)

	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}
}
