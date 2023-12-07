package lib

func CalculateWonHands(hands []Hand) []int {
	occurrenciesOfCards := make([]int, len(hands))

	for handIndex := range hands {
		occurrenciesOfCards[handIndex] = 1
	}

	for currentCardIndex := range occurrenciesOfCards {
		wonByHand := hands[currentCardIndex].WinningNumbers()

		for wonIndex := range wonByHand {
			occurrenciesOfCards[currentCardIndex+(wonIndex+1)] += occurrenciesOfCards[currentCardIndex]
		}
	}

	return occurrenciesOfCards
}
