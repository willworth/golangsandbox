package main

import (
	"example.com/example/9taxProject/cmdmanager" //alias!
	"example.com/example/9taxProject/prices"
)

func main() {

	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
