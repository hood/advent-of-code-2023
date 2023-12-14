package lib

import "slices"

func FloodFill(m *Map, c Coordinates, value rune, fillable []rune) *Map {
	// Get the current cell.
	current := m.At(c)

	// If it's already the correct color *what i f a wall* return.
	if current == value {
		return m
	}

	// Otherwise call the fill function which will fill in the existing image.
	fill(m, c, value, fillable)

	// Return the image once it is filled.
	return m
}

func fill(m *Map, c Coordinates, value rune, fillable []rune) {
	if m.IsOutOfBounds(c) {
		return
	}

	// If the current pixel is not which needs to be replaced
	if !slices.Contains(fillable, m.At(c)) {
		return
	}

	// Update the new color
	m.SetAt(c, value)

	// Fill in all four directions
	fill(m, Coordinates{X: c.X, Y: c.Y - 1}, value, fillable)
	fill(m, Coordinates{X: c.X, Y: c.Y + 1}, value, fillable)
	fill(m, Coordinates{X: c.X - 1, Y: c.Y}, value, fillable)
	fill(m, Coordinates{X: c.X + 1, Y: c.Y}, value, fillable)
}
