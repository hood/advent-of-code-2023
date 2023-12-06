package lib

func CountMinimumNeededGameCubesSet(gameCubesSets []CubesSet) CubesSet {
	minimum := CubesSet{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, partialCubesSet := range gameCubesSets {
		if partialCubesSet.Red > minimum.Red {
			minimum.Red = partialCubesSet.Red
		}

		if partialCubesSet.Green > minimum.Green {
			minimum.Green = partialCubesSet.Green
		}

		if partialCubesSet.Blue > minimum.Blue {
			minimum.Blue = partialCubesSet.Blue
		}
	}

	return minimum
}
