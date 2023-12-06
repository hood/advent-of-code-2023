package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseSeedsRanges(lines []string) []SeedsRange {
	seeds := []int{}

	for _, line := range lines {
		if strings.Contains(line, "seeds:") {
			seedsString := strings.Split(line, "seeds: ")[1]

			seedsStringArray := strings.Split(seedsString, " ")

			for _, seedString := range seedsStringArray {
				seeds = append(seeds, shared.ParseInteger(seedString))
			}
		}
	}

	ranges := []SeedsRange{}

	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, SeedsRange{Start: seeds[i], End: seeds[i] + seeds[i+1]})
	}

	return ranges
}
