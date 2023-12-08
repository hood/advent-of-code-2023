package lib

func FindInHashMapWithCompass(hashMap HashMap, c *Compass, target string) (bool, int) {
	level := 0
	current := hashMap["AAA"]
	direction := c.NextDirection()

	for {
		nextID := ""

		if direction == 'L' {
			nextID = current.L
		} else if direction == 'R' {
			nextID = current.R
		}

		if nextID == "ZZZ" {
			return true, level + 1
		}

		current = hashMap[nextID]

		direction = c.NextDirection()
		level++
	}
}

func FindInHashMapWithDoublePath(
	hashMap HashMap,
	c *Compass,
	startingPoints []string,
	isTarget func(v string) bool,
) (bool, int) {
	level := 0
	direction := c.NextDirection()

	cursors := []HashMapNode{}
	for _, startingPoint := range startingPoints {
		cursors = append(cursors, hashMap[startingPoint])
	}

	for {
		nextIDs := make([]string, len(startingPoints))

		if direction == 'L' {
			for i := range startingPoints {
				nextIDs[i] = cursors[i].L
			}
		} else if direction == 'R' {
			for i := range startingPoints {
				nextIDs[i] = cursors[i].R
			}
		}

		nodesAtEnd := 0
		for i := range nextIDs {
			if isTarget(nextIDs[i]) {
				nodesAtEnd++

				break
			}
		}
		if nodesAtEnd == len(nextIDs) {
			return true, level + 1
		}

		for i := range nextIDs {
			cursors[i] = hashMap[nextIDs[i]]
		}

		// println("Level", level, "direction", string(direction), "nextIDs", nextIDs)

		direction = c.NextDirection()
		level++
	}
}
