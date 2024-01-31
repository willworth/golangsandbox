package main

import (
	"fmt"

	filemanager "example.com/example/9taxProject/fileManager" //alias!
	"example.com/example/9taxProject/prices"
)

func main() {

	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
