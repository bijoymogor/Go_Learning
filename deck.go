package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck
//which is a slice of string

type deck []string

// d is direct reference call by reference similar to this
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"ace", "two", "three", "four"}

	// for i, suit := range cardSuits{
	// 	for j, value := range cardsValues{
	// 		cards = append(cards, value+" of "+ suit)
	// 	}
	// }

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// SubArrays
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error :-", error)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index, _ := range d {
		newPositon := r.Intn(len(d) - 1)

		d[index], d[newPositon] = d[newPositon], d[index]
	}
}
