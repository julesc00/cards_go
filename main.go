package main

import (
	"fmt"
)

var dice string

func main() {
	// Create a slice of strings
	cards := deck{
		"Ace of Spades",
		newCard(),
	}

	// Appending items to a slice
	cards = append(cards, "King of Spades")
	// Iterate over a slice, _ is the index that can also be used
	cards.print()

	dice = "dice"
	cardNumber := cardNumber()
	fmt.Println(cards)
	fmt.Println(dice)
	fmt.Println(cardNumber)
	printMessages()
}

func newCard() string {
	return "Five of Diamonds"
}

func cardNumber() int {
	return 8
}

func printMessages() {
	fmt.Println(len("Hello World!"))

	// This will return the byte 101 and not necessarily the value at index 1 which would be `e`
	fmt.Println("Hello, world"[1])
	fmt.Println("Hello" + "world!")
}
