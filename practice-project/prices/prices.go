package prices

import (
	"calculator/convertion"
	"calculator/iomanager"
	"fmt"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPricesJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPricesJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	prices, err := convertion.StringToFloat64(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadData()
	//errChan <- errors.New("asdasd")
	if err != nil {
		fmt.Println("Error loading data:", err)
		errChan <- err
		return
	}
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPricesJob := price + (price * job.TaxRate)
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPricesJob)
	}
	job.TaxIncludedPrices = results
	job.IOManager.WriteResult(job)
	doneChan <- true

}
