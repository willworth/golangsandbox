package main

import "fmt"

func main() {

	// card := "ace of spades"
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "five of diamonds"
}
