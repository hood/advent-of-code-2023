package lib

import "strings"

func FindStartingPoints(lines []string) []string {
	startingPoints := []string{}

	for _, line := range lines {
		if strings.Contains(line, "A = ") {
			nodeID := strings.Split(line, " = ")[0]
			startingPoints = append(startingPoints, nodeID)
		}
	}

	return startingPoints
}
