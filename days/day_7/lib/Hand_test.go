// package lib

// import "testing"

// func TestHandWithFourOfAKind(t *testing.T) {
// 	hand := &Hand{}

// 	input := []rune{'9', '9', '9', '9', 'K'}

// 	expected := int(1)

// 	for _, card := range input {
// 		hand.AddCard(card)
// 	}

// 	result := hand.Score()

// 	if result != expected {
// 		t.Errorf("Found %d, expected %d", result, expected)
// 	}
// }

// func TestHandWithFullHouse(t *testing.T) {
// 	hand := &Hand{}

// 	input := []rune{'9', '9', '9', 'K', 'K'}

// 	expected := int(10)

// 	for _, card := range input {
// 		hand.AddCard(card)
// 	}

// 	result := hand.Score()

// 	if result != expected {
// 		t.Errorf("Found %d, expected %d", result, expected)
// 	}
// }

package lib

import "testing"

func TestHandWithFourOfAKind(t *testing.T) {
	hand := NewHand(5)

	input := []rune{'9', '9', '9', '9', 'K'}

	expected := Score(1)

	for position, card := range input {
		println("Set card")
		hand.SetCard(position, Card(card))
	}

	result := hand.Eval()

	if result != expected {
		t.Errorf("Found %d, expected %d", result, expected)
	}
}
