package lib

import "strings"

// Given a line, returns the cubes set at the given index.
func ParseCubesSet(line string, cubesSetIndex int) CubesSet {
	cubesSet := CubesSet{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	// Stringified cube set at the given index.
	substring := strings.Split(line, ";")[cubesSetIndex]

	stringifiedCounts := strings.Split(substring, ",")

	for _, stringifiedCount := range stringifiedCounts {
		trimmedStringifiedCount := strings.TrimSpace(stringifiedCount)

		if strings.Contains(trimmedStringifiedCount, "red") {
			cubesSet.Red += ParseCubesCount(trimmedStringifiedCount)
		}

		if strings.Contains(trimmedStringifiedCount, "green") {
			cubesSet.Green += ParseCubesCount(trimmedStringifiedCount)
		}

		if strings.Contains(trimmedStringifiedCount, "blue") {
			cubesSet.Blue += ParseCubesCount(trimmedStringifiedCount)
		}
	}

	return cubesSet
}
