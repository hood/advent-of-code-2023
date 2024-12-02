package day_10

import (
	"adventofcode2023/days/day_10/lib"
	"adventofcode2023/days/shared"
	"testing"
)

// .┌────┐┌┐┌┐┌┐┌─┐....
// .│┌──┐││││││││┌┘....
// .││.┌┘││││││││└┐....
// ┌┘└┐└┐└┘└┘││└┘I└─┐..
// └──┘.└┐III└┘S┐┌─┐└┐.
// ....┌─┘II┌┐┌┘│└┐└┐└┐
// ....└┐I┌┐││└┐│I└┐└┐│
// .....│┌┘└┘│┌┘│┌┐│.└┘
// ....┌┘└─┐.││.││││...
// ....└───┘.└┘.└┘└┘...

func TestDay10Part2Test(t *testing.T) {
	lines := shared.ReadFile("./part_2_test_input.txt")

	m, start := lib.MapFromLines(lines)

	_, walked := m.FindFarthestTile(start)

	// check walked for duplicates
	for i := 0; i < len(walked)-1; i++ {
		println(string(walked[i].Value))
		for j := i + 1; j < len(walked); j++ {
			if walked[i].Coordinates.SameAs(walked[j].Coordinates) {
				t.Fatal("Duplicates found!")
			}
		}
	}

	println(string(walked[0].Value))
	println(string(walked[len(walked)-2].Value))

	walked[len(walked)-1].Value = 'F'

	result := lib.CalcPolygonArea2(walked)

	println("RESULT", result)

	shared.AssertEqual(t, 8, result)
}
