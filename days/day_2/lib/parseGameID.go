package lib

import (
	"strconv"
	"strings"
)

func ParseGameID(line string) int {
	parsedGameID, error := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], "Game ")[1])
	if error != nil {
		panic(error)
	}

	return parsedGameID
}
