package lib

import (
	"strconv"
	"strings"
)

// Counts the number of cubes in a string like "9 red".
func ParseCubesCount(stringifiedCount string) int {
	parsedCount, error := strconv.Atoi(strings.Split(stringifiedCount, " ")[0])
	if error != nil {
		panic(error)
	}

	return parsedCount
}
