package main

import (
	"fmt"
	"math"
)

func main() {
	var investment float64
	var interestRate float64
	var years int
	var inflationRate float64

	fmt.Print("Input investment: ")
	fmt.Scan(&investment)

	fmt.Print("Input interest rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Input years: ")
	fmt.Scan(&years)

	fmt.Print("Input inflation rate: ")
	fmt.Scan(&inflationRate)

	var futureValue = investment * math.Pow(1+interestRate/100, float64(years))
	var realFutureValue = futureValue / math.Pow(1+inflationRate/100, float64(years))

	fmt.Println("Total value: ", futureValue)
	fmt.Println("Total value after inflation: ", realFutureValue)
}
