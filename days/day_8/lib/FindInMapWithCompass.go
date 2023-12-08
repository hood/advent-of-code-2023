package lib

func FindInMapWithCompass(m *Map, c *Compass, target string) (bool, int) {
	level := 0
	current := m.Head
	direction := c.NextDirection()

	for current != nil {
		if current.ID == target {
			return true, level
		}

		// println(
		// 	"Current: ", current.ID,
		// 	"Level: ", level,
		// 	"Next: ", string(direction),
		// 	"Left: ", current.Left,
		// 	"Right: ", current.Right,
		// )

		if direction == 'L' {
			current = current.Left
		} else if direction == 'R' {
			current = current.Right
		}

		direction = c.NextDirection()
		level++
	}

	return false, -1
}

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
