package lib

import "testing"

func TestHandWithFourOfAKind(t *testing.T) {
	hand := &Hand{}

	input := []rune{'9', '9', '9', '9', 'K'}

	expected := uint64(1)

	for _, card := range input {
		hand.AddCard(card)
	}

	result := hand.Score()

	if result != expected {
		t.Errorf("Found %d, expected %d", result, expected)
	}
}

func TestHandWithFullHouse(t *testing.T) {
	hand := &Hand{}

	input := []rune{'9', '9', '9', 'K', 'K'}

	expected := uint64(10)

	for _, card := range input {
		hand.AddCard(card)
	}

	result := hand.Score()

	if result != expected {
		t.Errorf("Found %d, expected %d", result, expected)
	}
}
