package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, instead got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, insead got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")

	deck := newDeck()

	err := deck.saveToFile("_decktesting")

	if err != nil {
		t.Errorf("Unable to save _decktesting")
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) < 1 {
		t.Errorf("Expected deck length of at least 1, instead got %v", len(loadedDeck))
	}

	_ = os.Remove("_decktesting")
}
