package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestSequenceParse(t *testing.T) {
	input := "-3 -11 -26 -51 -85 -123 -145 -71 358 1848 6104 16888 41913 96030 206362 420524 820437 1547812 2855682 5217834 9561853"

	sequence := SequenceFromString(input)

	shared.AssertEqual(t, len(input), len(sequence.ValuesAsString()))
	shared.AssertEqual(t, input, sequence.ValuesAsString())
}
