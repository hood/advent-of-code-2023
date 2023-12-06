package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseSeeds(lines []string) []int {
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

	return seeds
}
