package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseRaces(lines []string) []Race {
	times := shared.ExtractIntegersFromString(strings.Split(lines[0], "Time: ")[1])
	distances := shared.ExtractIntegersFromString(strings.Split(lines[1], "Distance: ")[1])
	races := []Race{}

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			Time:     times[i].Value,
			Distance: distances[i].Value,
		})
	}

	return races
}
