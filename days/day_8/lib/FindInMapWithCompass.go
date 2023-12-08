package lib

func FindInMapWithCompass(m *Map, c *Compass, target string) (bool, int) {
	level := 0

	current := m.Head

	for current != nil {
		if current.ID == target {
			return true, level
		}

		next := c.NextDirection()

		if next == 'L' {
			current = current.Left
		} else if next == 'R' {
			current = current.Right
		}

		level++
	}

	return false, -1
}
