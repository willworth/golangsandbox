package main

import (
	"fmt"

	filemanager "example.com/example/9taxProject/fileManager"
	"example.com/example/9taxProject/prices"
)

func main() {

	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}
	// wait for every chan to emit 1 val
	for _, doneChan := range doneChans {
		<-doneChan

	}
}
