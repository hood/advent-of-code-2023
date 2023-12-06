package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseMap(mapString string) Map {
	indices := strings.Split(mapString, " ")

	return Map{
		DestinationRangeStart: shared.ParseInteger(indices[0]),
		SourceRangeStart:      shared.ParseInteger(indices[1]),
		RangeLength:           shared.ParseInteger(indices[2]),
	}
}
