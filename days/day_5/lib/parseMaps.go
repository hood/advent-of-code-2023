package lib

import "adventofcode2023/days/shared"

func ParseMaps(mapStrings []string) *shared.BinarySearchTree[Map] {
	maps := shared.NewBinarySearchTree[Map]()

	for _, mapString := range mapStrings {
		maps.Insert(ParseMap(mapString))
	}

	return maps
}
