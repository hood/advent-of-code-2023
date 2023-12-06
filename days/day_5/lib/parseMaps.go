package lib

func ParseMaps(mapStrings []string) []Map {
	maps := make([]Map, len(mapStrings))

	for index, mapString := range mapStrings {
		maps[index] = ParseMap(mapString)
	}

	return maps
}
