package main

import (
	"fmt"
)

func main() {
	fmt.Println("func package")

	numbers := []int{1, 2, 3, 4}
	// doubled := doubleNumbers(&numbers)
	doubled := transformNumbers(&numbers, double) // pasing a func as a param
	tripled := transformNumbers(&numbers, triple) // pasing a func as a param
	fmt.Println("doubled", doubled)
	fmt.Println("tripled", tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}
func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
