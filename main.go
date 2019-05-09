package main

func main() {
	cards := newDeck()

	_ = cards.saveToFile("my_cards")
}
