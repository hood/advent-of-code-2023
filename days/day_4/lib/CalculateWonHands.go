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

// func CalculateWonHands(hands []Hand) []int {
// 	// results := make([]int, len(hands)*10)

// 	for index := range hands {
// 		wonByHand := hands[index].WinningNumbers()

// 		println("Card", index+1, "won", len(wonByHand), "hands.")

// 		for winIndex := range wonByHand {
// 			wonCardIndex := winIndex + (index + 1)

// 			println("About to iterate", max(results[wonCardIndex], 1), "times for card", wonCardIndex+1)

// 			instancesOfHand := results[wonCardIndex] + 1
// 			for i := 0; i < instancesOfHand; i++ {
// 				results[wonCardIndex] += len(hands[wonCardIndex].WinningNumbers())
// 			}
// 		}

// 		// 	results[index+(winIndex+1)] += len(hands[index+(winIndex+1)].WinningNumbers())
// 	}

// 	return results
// }

// func GenerateWonHandsIndices(startingFrom int, hands []Hand) int {
// 	wonByHand := [200]int{}

// 	won := hands[startingFrom].WinningNumbers()

// 	wonByHand[startingFrom] = len(won)

// 	println("Will check won cards for hand ", startingFrom+1, "of", len(hands)-1, "for", len(won), "times.")
// 	for index := range won {
// 		wonByHand[startingFrom+(index+1)] += GenerateWonHandsIndices(startingFrom+(index+1), hands)
// 	}

// 	result := 0

// 	for _, won := range wonByHand {
// 		result += won
// 	}

// 	return result
// }

// func GenerateWonHandsIndices(startingFrom int, hands []Hand) int {

// 	result := 0

// 	won := hands[startingFrom].WinningNumbers()
// 	result += len(won)

// 	for index := range won {
// 		result += GenerateWonHandsIndices(startingFrom+(index+1), hands)
// 	}

// 	return result
// }

// func GenerateWonHandsIndices(startingFrom int, hands []Hand) int {
// 	wonHands := 0

// 	winningNumbers := hands[startingFrom].WinningNumbers()

// 	if startingFrom <= len(hands) {
// 		for index := range winningNumbers {
// 			wonHands += GenerateWonHandsIndices(startingFrom+(index+1), hands)
// 		}
// 	}

// 	return wonHands + len(winningNumbers)
// }

// func GenerateWonHandsIndices(startingFrom int, hands []Hand) []int {
// 	wonHandsIndices := []int{}

// 	if startingFrom == len(hands)-1 {
// 		return wonHandsIndices
// 	}

// 	winningNumbers := hands[startingFrom].WinningNumbers()

// 	for index := range winningNumbers {
// 		wonHandsIndices = append(wonHandsIndices, GenerateWonHandsIndices(startingFrom+(index+1), hands)...)
// 	}

// 	println("wonHandsIndices", "->", len(wonHandsIndices))

// 	return wonHandsIndices
// }
