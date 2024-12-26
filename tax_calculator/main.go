package main

import (
	"fmt"
	"tax-calculator/filemanager"
	"tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.25, 0.35}

	for _, taxRate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_tax_%.0f.json", taxRate*100))

		// cmdManager := cmdmanager.New()
		taxJob := prices.New(taxRate, fileManager)
		err := taxJob.Process()
		if err != nil {
			fmt.Println("Cannot process calculating tax, error: ", err)
			return
		}
	}

}
