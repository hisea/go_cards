package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d %s\n", i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(b), ","))
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		j := r.Intn(len(d))
		d[j], d[i] = d[i], d[j]
	}
}
