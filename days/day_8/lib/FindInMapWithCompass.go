package lib

func FindInHashMapWithCompass(
	hashMap HashMap,
	c *Compass,
	startingPoint string,
	isTarget func(v string) bool,
) (bool, int) {
	c.Reset()

	level := 0
	current := hashMap[startingPoint]
	direction := c.NextDirection()

	for {
		nextID := ""

		if direction == 'L' {
			nextID = current.L
		} else if direction == 'R' {
			nextID = current.R
		}

		if isTarget(nextID) {
			return true, level + 1
		}

		current = hashMap[nextID]

		direction = c.NextDirection()
		level++
	}
}
