package lib

import "adventofcode2023/days/shared"

func GetPairs(rangeMap Map) [][]int {
	pairs := [][]int{}

	for step := range shared.Times(rangeMap.RangeLength) {
		pairs = append(pairs, []int{rangeMap.SourceRangeStart + step, rangeMap.DestinationRangeStart + step})
	}

	return pairs
}
