package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 15.55, 20.0}

	prices[3] = 1
	fmt.Println(prices)

	featurePrices := prices[1:]
	highlightedPrices := featurePrices[:1]
	fmt.Println(featurePrices)
	fmt.Println(highlightedPrices)
}