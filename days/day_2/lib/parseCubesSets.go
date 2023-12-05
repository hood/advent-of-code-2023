package lib

import (
	"strings"
)

func ParseCubesSets(line string) []CubesSet {
	cubesSetsCount := CountCubesSets(line)
	cubesSets := make([]CubesSet, cubesSetsCount)

	for i := 0; i < cubesSetsCount; i++ {
		cubesSets[i] = ParseCubesSet(strings.Split(line, ":")[1], i)
	}

	return cubesSets
}
