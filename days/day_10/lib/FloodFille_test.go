package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestFloodFill(t *testing.T) {
	input := []string{
		"....",
		".FT.",
		".LJ.",
		"....",
	}

	m, _ := MapFromLines(input)

	FloodFill(&m, Coordinates{X: 0, Y: 0}, 'X')

	expected := []string{
		"XXXX",
		"XFTX",
		"XLJX",
		"XXXX",
	}

	for rowIndex := range expected {
		for columnIndex := range expected[rowIndex] {
			if input[rowIndex][columnIndex] == '.' {
				shared.AssertEqual(t, 'X', m.At(Coordinates{
					X: columnIndex,
					Y: rowIndex,
				}), m.ToLines()...)
			}
		}
	}
}
