package main

func main() {
	cards := newDeck()
	cards.saveToFile("deck.txt")
	newCards := newDeckFromFile("deck.txt")
	newCards.print()
}
