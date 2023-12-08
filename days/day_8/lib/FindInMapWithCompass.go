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
