package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputNumber float64
	fmt.Scan(&inputNumber)

	writeInputNumberToFile(inputNumber)

	fmt.Println(getBalanceFromFile())
}

func writeInputNumberToFile(inputNumber float64) {
	inputNumberText := fmt.Sprint(inputNumber)
	os.WriteFile("filename.txt", []byte(inputNumberText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile("filename.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	
	return balance
}
