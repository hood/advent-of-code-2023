package lib

import "slices"

type Hand struct {
	ID          int
	OwnCard     Card
	WinningCard Card
}

func (hand *Hand) WinningNumbers() []int {
	winningNumbers := []int{}

	for ownNumber := range hand.OwnCard.Numbers {
		if slices.Contains(hand.WinningCard.Numbers, ownNumber) {
			winningNumbers = append(winningNumbers, ownNumber)
		}
	}

	return winningNumbers
}
