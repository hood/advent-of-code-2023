package lib

func IsGameFeasible(gameCubesSets []CubesSet, availableCubesSet CubesSet) bool {
	for _, partialCubesSet := range gameCubesSets {
		if partialCubesSet.Red > availableCubesSet.Red ||
			partialCubesSet.Green > availableCubesSet.Green ||
			partialCubesSet.Blue > availableCubesSet.Blue {
			return false
		}
	}

	return true
}
