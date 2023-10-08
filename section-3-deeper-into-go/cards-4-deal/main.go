package main

func main() {
	cards := newDeck()

	hand, remaingCards := deal(cards, 5)

	hand.print()
	remaingCards.print()
}
