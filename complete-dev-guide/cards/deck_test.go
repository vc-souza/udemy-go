package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktest"

	os.Remove(filename)

	original := newDeck()

	original.saveToFile(filename)

	loaded := newDeckFromFile(filename)

	if len(loaded) != len(original) {
		t.Errorf("Expected %v got %v", len(original), len(loaded))
	}

	os.Remove(filename)
}
