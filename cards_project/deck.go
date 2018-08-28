package main

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardNum := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuite {
		for _, num := range cardNum {
			cards = append(cards, num + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) writeToFile(filename string) error {
	strDeck := d.convertSliceToString()
	byteSlice := []byte(strDeck)
	return ioutil.WriteFile(filename, byteSlice, 0666)
}

func (d deck) convertSliceToString() string {
	// Convert deck to slice of strings
	strDeck := []string(d)
	// Join slice of string with ", "
	return strings.Join(strDeck, ", ")
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	strData := string(byteSlice)
	stringSlice := strings.Split(strData, ", ")
	return deck(stringSlice)
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap, same a python
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}