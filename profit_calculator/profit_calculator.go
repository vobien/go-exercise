package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Input revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expense, err := getUserInput("Input expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Input tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancialMetrics(revenue, expense, taxRate)
	writeResult(ebt, profit, ratio)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

}

func writeResult(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(result), 0644)
}

func calculateFinancialMetrics(revenue, expense, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expense
	var profit = ebt * (1 - taxRate/100)
	var ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(labelText string) (float64, error) {
	var userInput float64
	fmt.Print(labelText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		return 0, errors.New("invalid input. Please enter a positive number")
	}

	return userInput, nil
}
