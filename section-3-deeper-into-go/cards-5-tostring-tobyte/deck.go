package main

import (
	"fmt"
	"os"
	"strings"
)

// we will create a custom type named deck
// deck will be a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two", "One"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits { // we want to add the value and the suits together each
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d { // means this loop will read as long as the range of d
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // return 2 deck type value. one will be the slice[start_of_slice_value_not_included:handsize] and the other will be the slice[handsize:end_of_slice_value_included]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // standard library funtion to convert a slice of strings to a single string saparated by ","
}

func (d deck) toByte() []byte {
	return []byte(d.toString()) // we are converting the string from the toString function and turning it into a slice of bytes so that we can save it to our hard drive
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByte(), 0666)
}
