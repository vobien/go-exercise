package main

import (
	"tax-calculator/filemanager"
	"tax-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.1, 0.25, 0.35}

	for _, taxRate := range taxRates {
		taxJob := prices.New(taxRate)
		taxJob.Calculate()

		path := fmt.Sprintf("result_tax_%.0f.json", taxRate * 100)
		err := filemanager.WriteJSON(path, taxJob)
		if err != nil {
			fmt.Printf("Error writing JSON to path %s, error: %s", path, err)
			return
		}
	}

}
