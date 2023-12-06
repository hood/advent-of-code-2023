package lib

func CalculateWonHands(hands []Hand) []int {
	occurrenciesOfCards := make([]int, len(hands))

	for handIndex := range hands {
		occurrenciesOfCards[handIndex] = 1
	}

	for currentCardIndex := range occurrenciesOfCards {
		wonByHand := hands[currentCardIndex].WinningNumbers()

		for occurrencyOfCurrentCard := 0; occurrencyOfCurrentCard < occurrenciesOfCards[currentCardIndex]; occurrencyOfCurrentCard++ {
			for wonIndex := range wonByHand {
				occurrenciesOfCards[currentCardIndex+(wonIndex+1)] += 1
			}
		}
	}

	return occurrenciesOfCards
}
