package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseRaces(lines []string) []Race {
	times := []int{}
	distances := []int{}
	races := []Race{}

	for _, char := range strings.Split(lines[0], "Time: ")[1] {
		for _, stringifiedTime := range strings.Split(string(char), " ") {
			if stringifiedTime == "" {
				continue
			}

			times = append(times, shared.ParseInteger(stringifiedTime))
		}
	}

	for _, char := range strings.Split(lines[1], "Distance: ")[1] {
		for _, stringifiedDistance := range strings.Split(string(char), " ") {
			if stringifiedDistance == "" {
				continue
			}

			distances = append(distances, shared.ParseInteger(stringifiedDistance))
		}
	}

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			Time:     times[i],
			Distance: distances[i],
		})
	}

	return races
}
