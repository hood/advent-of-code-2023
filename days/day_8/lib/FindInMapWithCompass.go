package lib

func FindInMapWithCompass(m *Map, c *Compass, target string) (bool, int) {
	level := 0
	current := m.Head
	direction := c.NextDirection()

	for current != nil {
		if current.ID == target {
			return true, level
		}

		println(
			"Current: ", current.ID,
			"Level: ", level,
			"Next: ", string(direction),
			"Left: ", current.Left.ID,
			"Right: ", current.Right.ID,
		)

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
