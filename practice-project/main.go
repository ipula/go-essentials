package main

import (
	"calculator/filemanager"
	"calculator/prices"
	"fmt"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		go pricesJob.Process(doneChans[i], errorChans[i])
		// if err != nil {
		// 	fmt.Println("Error processing prices:", err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Printf("Error processing prices with tax rate %.2f%%: %v\n", taxRates[index]*100, err)
			}
		case <-doneChans[index]:
			fmt.Printf("Finished processing prices with tax rate %.2f%%\n", taxRates[index]*100)
		}
	}
}
