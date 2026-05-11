package main

import (
	"calculator/filemanager"
	"calculator/prices"
	"fmt"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		err := pricesJob.Process()
		if err != nil {
			fmt.Println("Error processing prices:", err)
		}
	}
}
