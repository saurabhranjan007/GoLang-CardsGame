package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	// instantiate the given func
	d := newDeck()

	// Test101: Check for the length of the created Deck.
	if len(d) != 20 {
		// if length don't match => notify t handler
		// Injecting the value of len(d) into the t hanlder
		t.Errorf("Expected deck length of 20, got %v", len(d))
	}

	// Testing the first card value
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}
}

func TestSaveTodeckAndNewDeckFile(t *testing.T) {

	fmt.Println("First time removal ..")

	// dealing with file sys always remove any test file available
	os.Remove("_decktesting")

	// instantiating the fucntions that are to be tested
	deck := newDeck()
	deck.savetToFile("_decktesting")

	// Load Test
	loadedDeck := newDeckFromFile("_decktesting")

	fmt.Println("Deck initialised and loaded")

	// Test Cases ⬇️
	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
