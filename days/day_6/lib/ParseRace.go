package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

// ParseRace is the same as parse races, but spaces between numbers are ignored
// and everything counts as the same, big race.
func ParseRace(lines []string) Race {
	time := shared.ExtractIntegersFromString(
		strings.ReplaceAll(
			strings.Split(lines[0], "Time: ")[1],
			" ",
			"",
		),
	)[0]
	distance := shared.ExtractIntegersFromString(
		strings.ReplaceAll(
			strings.Split(lines[1], "Distance: ")[1],
			" ",
			"",
		),
	)[0]

	return Race{
		Time:     time.Value,
		Distance: distance.Value,
	}
}
