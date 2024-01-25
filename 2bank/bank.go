package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	"example.com/inv-calc/2bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	fmt.Println("Welcome to GO Bank")
	fmt.Println("Call us", randomdata.PhoneNumber())

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----")
		// panic("I can't go on!")
	}

	for {

		presentOptions()

		var choice int
		fmt.Print("Please choose an option:")
		fmt.Scan(&choice)

		fmt.Println("Your choice", choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64 // only visible in this branch
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64 // only visible in this branch
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough!.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else {
			fmt.Println("Bye")
			break
		}
	}
	fmt.Println("Thanks for using Go Bank!")
}
