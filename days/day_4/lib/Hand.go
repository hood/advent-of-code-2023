package lib

import "slices"

type Hand struct {
	ID          int
	OwnCard     Card
	WinningCard Card
}

func (hand *Hand) WinningNumbers() []int {
	winningNumbers := []int{}

	for _, ownNumber := range hand.OwnCard.Numbers {
		if slices.Contains(hand.WinningCard.Numbers, ownNumber) {
			winningNumbers = append(winningNumbers, ownNumber)
		}
	}

	return winningNumbers
}

func (hand *Hand) Points() int {
	points := 0

	for index := range hand.WinningNumbers() {
		if index == 0 {
			points++
		} else {
			points += points
		}
	}

	return points
}
