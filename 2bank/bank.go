package main

import "fmt"
func main()	{
	fmt.Println("Welcome to GO Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("4. Exit")


	var choice int
	fmt.Print("Please choose an option:")
	fmt.Scan(&choice)

	fmt.Println("Your choice", choice)
} 