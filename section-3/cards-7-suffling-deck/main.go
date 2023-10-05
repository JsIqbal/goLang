package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print() // the suffled deck is always same. we will talk about it in cards-9-random-number
}
