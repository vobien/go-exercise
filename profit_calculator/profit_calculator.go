package main

import "fmt"

func main() {

	revenue := getUserInput("Input revenue: ")
	expense := getUserInput("Input expenses: ")
	taxRate := getUserInput("Input tax rate: ")

	ebt, profit, ratio := calculateFinancialMetrics(revenue, expense, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

}

func calculateFinancialMetrics(revenue, expense, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expense
	var profit = ebt * (1 - taxRate/100)
	var ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(labelText string) float64 {
	var userInput float64
	fmt.Print(labelText)
	fmt.Scan(&userInput)
	return userInput
}
