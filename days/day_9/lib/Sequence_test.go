package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestSequenceFromString(t *testing.T) {
	input := "-3 -11 -26 -51 -85 -123 -145 -71 358 1848 6104 16888 41913 96030 206362 420524 820437 1547812 2855682 5217834 9561853"

	sequence := SequenceFromString(input)

	shared.AssertEqual(t, len(input), len(sequence.ValuesAsString()))
	shared.AssertEqual(t, input, sequence.ValuesAsString())
}

func TestSequenceStepSizes(t *testing.T) {
	input := "1 3 6 10 15 21"

	expected := []string{
		"1 3 6 10 15 21 28",
		"2 3 4 5 6 7",
		"1 1 1 1 1",
	}

	sequence := SequenceFromString(input)

	shared.AssertEqual(t, len(expected), len(sequence.StepSizesAsString()))
	shared.AssertEqual(t, expected[0], sequence.StepSizesAsString()[0])
}

func TestSequenceFinalValue(t *testing.T) {
	input := "1 3 6 10 15 21"

	sequence := SequenceFromString(input)

	shared.AssertEqual(t, 28, sequence.FinalValue())
}

func TestSequenceFromString2(t *testing.T) {
	input := "8 6 1 -7 -18 -32 -49 -69 -92 -118 -147 -179 -214 -252 -293 -337 -384 -434 -487 -543 -602"

	sequence := SequenceFromString(input)

	shared.AssertEqual(t, len(input), len(sequence.ValuesAsString()))
	shared.AssertEqual(t, input, sequence.ValuesAsString())
}

func TestSequenceFromStringFull(t *testing.T) {
	lines := shared.ReadFile("../input.txt")

	for _, line := range lines {
		sequence := SequenceFromString(line)

		shared.AssertEqual(t, len(line), len(sequence.ValuesAsString()))
		shared.AssertEqual(t, line, sequence.ValuesAsString())
	}
}

func TestSequenceFinalValueEdgeCase(t *testing.T) {
	line := "4 17 47 106 214 398 691 1131 1760 2623 3767 5240 7090 9364 12107 15361 19164 23549 28543 34166 40430"
	sequence := SequenceFromString(line)

	expected := 47338

	shared.AssertEqual(t, line, sequence.ValuesAsString())
	shared.AssertEqual(t, expected, sequence.FinalValue())
}
