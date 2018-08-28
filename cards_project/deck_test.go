package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Excepted 1st in deck to be 'Ace of Spades', got %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Excepted last in deck to be 'Four of Hearts', got %v", d[len(d)-1])
	}
}
