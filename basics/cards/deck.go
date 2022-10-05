package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type
type deck []string

// Create new deck
func newDeck() deck {

	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print a deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Split deck into two
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Change type from deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Create a file with the deck content
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Create a trully random seed every time the program run using the time
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new definition of the rand package but using our seed
	r := rand.New(source)

	for i := range d {
		// Instead of using the classic rand library we are using our new rand definition
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}
