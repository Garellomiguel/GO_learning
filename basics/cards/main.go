package main

func main() {

	// cards := newDeck()

	// hand, restOfDeck := deal(cards, 3)
	// hand.print()
	// restOfDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
}
