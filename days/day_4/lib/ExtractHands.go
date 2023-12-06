package lib

import (
	"strconv"
	"strings"
)

func ExtractHands(lines []string) []Hand {
	hands := []Hand{}

	for _, line := range lines {
		hand := Hand{}

		title := strings.Split(line, ":")[0]
		cardsStrings := strings.Split(line, ":")[1]

		id := strings.Split(title, "Card")[1]

		parsedID, error := strconv.Atoi(strings.Trim(id, " "))
		if error != nil {
			panic(error)
		}

		hand.ID = parsedID

		singleCardsStrings := strings.Split(cardsStrings, "|")

		hand.OwnCard = ExtractCard(singleCardsStrings[0])
		hand.WinningCard = ExtractCard(singleCardsStrings[1])

		hands = append(hands, hand)
	}

	return hands
}
