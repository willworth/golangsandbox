package main

import "fmt"
func main()	{


	var accountBalance = 1000.0

	fmt.Println("Welcome to GO Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("4. Exit")


	var choice int
	fmt.Print("Please choose an option:")
	fmt.Scan(&choice)

	fmt.Println("Your choice", choice)


	if choice == 1 {
fmt.Println("Your balance is", accountBalance)
	} else if choice ==2 {
		fmt.Print("How much do you want to deposit? ")
		var depositAmount float64  // only visible in this branch
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your new balance is", accountBalance)
	}else if choice ==3 {
		fmt.Print("How much do you want to withdraw? ")
		var withdrawAmount float64  // only visible in this branch
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Your new balance is", accountBalance)

	}
} 