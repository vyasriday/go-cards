package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Create a new type of deck which is a slice of string

type deck []string

// create a function that returns a new deck of cards

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// create a new function that takes a deck and loops over it
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) {

	cards := d.toString()

	err := ioutil.WriteFile(filename, []byte(cards), 0644)

	if err != nil {
		log.Fatal(err)
	}

}

func newDeckFromFile(filename string) deck {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return (toDeck(string(data)))

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toDeck(stringData string) deck {
	return deck(strings.Split(stringData, ","))
}
