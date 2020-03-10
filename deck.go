package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0664)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(fileName string) deck {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}
	return toDeck(data)

}

func toDeck(data []byte) deck {
	return deck(strings.Split(string(data), ","))
}
