package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	prices := make([]float64, 4, 5)

	prices[2] = 15.55
	prices[3] = 20.0

	prices = append(prices, 99.99)
	prices = append(prices, 99.99)
	fmt.Println(prices)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6

	courseRatings.output()

	//for range prices {}
	for index, value := range prices {
		fmt.Println(index, value)
	}
	
	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
