package lib

import "strings"

func GetMapsGroup(lines []string, title string) []string {
	startIndex, endIndex := 0, 0

	for index, line := range lines {
		if strings.Contains(line, title) {
			startIndex = index + 1

			break
		}
	}

	for index := startIndex; index < len(lines); index++ {
		if len(lines[index]) == 0 {
			endIndex = index

			break
		}
	}

	return lines[startIndex:endIndex]
}
