package lib

import "strings"

// Counts the number of cubes sets in a line.
func CountCubesSets(line string) int {
	return strings.Count(line, ";") + 1
}
