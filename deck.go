package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// New typpe of Deck, which is a slice of strings
type deck []string

// newDeck function of

// receiver to loop thro cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	// Create two slices one for available card suite names
	// And another for probable card value
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	// Creating the Deck Of Cards
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
			// fmt.Println("Card-", value+" of "+suite)
		}
	}

	return cards

}

// Deal func; will be called with the existing pack of cards "d deck"
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// UTIL FUNCTIONS turns deck into string (takes the whole of deck d, and converts it into string and then return it)
func (d deck) toString() string {
	// to convert data from one format into another, below is the syntax we use
	// []string/byte(data) => here converting the data into byte (byte slice)+
	// []string(d)
	convertedData := strings.Join([]string(d), ",")
	fmt.Println(convertedData)
	return convertedData
}

// save file in bytes to hard-drive
func (d deck) savetToFile(filename string) error {

	// using io/ioutils to save the file to drive.. 066 permission means anyone can read/write
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// get a deck from the saved file

func newDeckFromFile(filename string) deck {

	byte_slice, err := ioutil.ReadFile(filename)

	// two ways we could handle the error part here in our program
	// 1. Log the error and return a call to newDeck()
	// 2. Log the error and entirely quit the program

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// byteToStr := string(byte_slice)
	// fmt.Println(byteToStr)

	str := strings.Split(string(byte_slice), ",")
	fmt.Println(str)
	return deck(str)
}

// Randomize the Deck
func (d deck) shuffle() {

	// get radom source to generate randon int which shuffles the deck
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	fmt.Println("Spurce: ", source)
	fmt.Println("Random: ", r)

	for index := range d {
		// generating random num ⬇️
		// newPosition := rand.Intn(len(d) - 1)

		newPosition := r.Intn(len(d) - 1)
		fmt.Println("New Position: ", newPosition)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
