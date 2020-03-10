package main

import "fmt"

func main() {

	cards := []string{"Ace of Hearts", "King of Club"}
	cards = append(cards, newCard())
	var i int
	var card string
	for i, card = range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
