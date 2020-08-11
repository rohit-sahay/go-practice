package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamond" {
		t.Errorf("Expected last card of Four of Diamond, but got %v,", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "_deckTesting"
	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove(testFile)
}
