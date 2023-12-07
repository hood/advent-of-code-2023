package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseRaces(lines []string) []Race {
	times := []int{}
	durations := []int{}
	races := []Race{}

	for _, char := range strings.Split(lines[0], "Time: ")[1] {
		for _, stringifiedTime := range strings.Split(string(char), " ") {
			times = append(times, shared.ParseInteger(stringifiedTime))
		}
	}

	for _, char := range strings.Split(lines[0], "Duration: ")[1] {
		for _, stringifiedDuration := range strings.Split(string(char), " ") {
			durations = append(durations, shared.ParseInteger(stringifiedDuration))
		}
	}

	for i := 0; i < len(times); i++ {
		races[i] = Race{
			Time:     times[i],
			Duration: durations[i],
		}
	}

	return races
}
