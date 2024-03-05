package main

import "fmt"

func main() {
	prices := []float64{14.2, 52.6} // will autogrow on demand
	fmt.Println(prices)

	// you could assign to prices
	updatedPrices := append(prices, 5.99)
	fmt.Println(prices)
	fmt.Println(updatedPrices)

	// no remove function.  just slice to a new var
}

// func main() {

// 	var productNames [4]string
// 	prices := [4]float64{23.3, 45.3, 63.5, 77.2}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2]) //go is zero indexed

// 	// starting and excluding end ie index 3 is not included
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// 	// https://go.dev/tour/moretypes/7

// 	// if you edit a slice, it edits the original
// 	// len = quantity in an arr
// 	// cap = capacity

// }
