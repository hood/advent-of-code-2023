package lib

func FloodFill(m *Map, c Coordinates, value rune) *Map {
	// Get the current cell.
	current := m.At(c)

	// If it's already the correct color *what i f a wall* return.
	if current == value {
		return m
	}

	// Otherwise call the fill function which will fill in the existing image.
	fill(m, c, value, current)

	// Return the image once it is filled.
	return m
}

func fill(m *Map, c Coordinates, value rune, current rune) {
	// If row is less than 0
	if c.Y < 0 {
		return
	}

	// If column is less than 0
	if c.X < 0 {
		return
	}

	// If row is greater than image length
	if c.Y > len(*m)-1 {
		return
	}

	// If column is greater than image length
	if c.X > len((*m)[0])-1 {
		return
	}

	// If the current pixel is not which needs to be replaced
	if m.At(c) != current {
		return
	}

	// Update the new color
	m.SetAt(c, value)

	// Fill in all four directions
	// Fill Prev row
	fill(m, Coordinates{X: c.X, Y: c.Y - 1}, value, current)

	// Fill Next row
	fill(m, Coordinates{X: c.X, Y: c.Y + 1}, value, current)

	// Fill Prev col
	fill(m, Coordinates{X: c.X - 1, Y: c.Y}, value, current)

	// Fill next col
	fill(m, Coordinates{X: c.X + 1, Y: c.Y}, value, current)
}
