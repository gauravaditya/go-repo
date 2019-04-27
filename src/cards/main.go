package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards")

	cards := newDeckFromFile("mycards")
	fmt.Println(cards)
	cards.shuffle()
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
