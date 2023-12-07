package lib

import "adventofcode2023/days/shared"

// FindWaysToWinRace finds all the possible button pressing durations needed
// to win a race. Pushing a button for N milliseconds will result in the boat
// moving at a speed of N millimeter per millisecond for the duration of the
// remaining time. The results are returned in order of duration.
func FindWaysToWinRace(race *Race) []int {
	results := []int{}
	boat := Boat{}

	for _, buttonPressDuration := range shared.IntegersInRange(1, race.Time) {
		boat.SetSpeed(buttonPressDuration)

		distanceTraveledInRemainingRaceTime := boat.DistanceTraveledInTime(race.Time - buttonPressDuration)

		if distanceTraveledInRemainingRaceTime >= race.Distance {
			results = append(results, buttonPressDuration)
		}
	}

	return results
}
