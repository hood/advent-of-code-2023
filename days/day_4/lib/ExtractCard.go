package lib

import "adventofcode2023/days/shared"

func ExtractCard(s string) Card {
	card := Card{}

	values := shared.ExtractIntegersFromString(s)

	for _, value := range values {
		card.Numbers = append(card.Numbers, value.Value)
	}

	return card
}
