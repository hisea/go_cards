package main

type deck []string

func newDeck() deck {
	cardSuites := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	d := deck{}
	for _, s := range cardSuites {
		for _, v := range cardValues {
			d = append(d, v+" of "+s)
		}
	}
	return d
}
