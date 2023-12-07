package lib

import "testing"

func TestHandWithFourOfAKind(t *testing.T) {
	hand := &Hand{}

	input := []rune{'9', '9', '9', '9', 'K'}

	for _, card := range input {
		hand.AddCard(card)
	}

	result := hand.NumberScore()

	if result <= 0 {
		t.Errorf("Found %d, expected result > 0", result)
	}
}

func TestHandWithFullHouse(t *testing.T) {
	hand := &Hand{}

	input := []rune{'9', '9', '9', 'K', 'K'}

	for _, card := range input {
		hand.AddCard(card)
	}

	result := hand.NumberScore()

	if result < 166 {
		t.Errorf("Found %d, expected result > 166", result)
	}
}

func TestHandWithThreeOfAKind(t *testing.T) {
	hand := &Hand{}

	input := []rune{'9', '9', '9', 'K', '7'}

	for _, card := range input {
		hand.AddCard(card)
	}

	result := hand.NumberScore()

	if result < 1609 {
		t.Errorf("Found %d, expected result > 1609", result)
	}
}
