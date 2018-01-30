package main

func main() {
	d := newDeck()
	d.shuffle()
	d.print()
	// d.saveToFile("my_cards")
}
