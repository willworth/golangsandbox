package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio, err := calculateFinancials(revenue, expenses, taxRate)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

	err = storeResults(ebt, profit, ratio)
	if err != nil {
		fmt.Println(err)
	}
}

func storeResults(ebt, profit, ratio float64) error {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	return os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64, error) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	if profit == 0 {
		return 0, 0, 0, errors.New("Profit is zero, can't calculate ratio")
	}
	ratio := ebt / profit
	return ebt, profit, ratio, nil
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, err
	}
	if userInput < 0 {
		return 0, errors.New("Value must be non-negative")
	}
	return userInput, nil
}
