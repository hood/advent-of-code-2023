package lib

import (
	"testing"
)

func TestHandWithFourOfAKind(t *testing.T) {
	hand := NewHand()

	input := []rune{'9', '9', '9', '9', 'K'}

	for _, card := range input {
		hand.AddCard(card)
	}

	expected := FourOfAKind
	result := hand.NumberScore()

	if result != expected {
		t.Errorf("Found %d, expected %v", result, expected)
	}
}

func TestHandWithFullHouse(t *testing.T) {
	hand := NewHand()

	input := []rune{'9', '9', '9', 'K', 'K'}

	for _, card := range input {
		hand.AddCard(card)
	}

	expected := FullHouse
	result := hand.NumberScore()

	if result != expected {
		t.Errorf("Found %d, expected %v", result, expected)
	}
}

func TestHandWithThreeOfAKind(t *testing.T) {
	hand := NewHand()

	input := []rune{'9', '9', '9', 'K', '7'}

	for _, card := range input {
		hand.AddCard(card)
	}

	expected := ThreeOfAKind
	result := hand.NumberScore()

	if result != expected {
		t.Errorf("Found %d, expected %v", result, expected)
	}
}
