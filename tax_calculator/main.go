package main

import (
	"fmt"
	"tax-calculator/filemanager"
	"tax-calculator/prices"
	"time"
)

func main() {
	taxRates := []float64{0, 0.1, 0.25, 0.35}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	startTime := time.Now()
	fmt.Println("Start processing tax at ", startTime)
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_tax_%.0f.json", taxRate*100))

		// cmdManager := cmdmanager.New()
		taxJob := prices.New(taxRate, fileManager)

		go taxJob.Process(doneChans[index], errorChans[index])

		fmt.Println("Complete processing tax ", taxRate)
	}

	for id, _ := range taxRates {
		select {
		case err := <-errorChans[id]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[id]:
			fmt.Println("Complete processing tax ", taxRates[id])
		}
	}

	endTime := time.Now()
	fmt.Println("Finish processing tax at ", endTime)
	fmt.Println("Total time processing tax: ", endTime.Sub(startTime))
}
