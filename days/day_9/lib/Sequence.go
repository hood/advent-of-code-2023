package lib

import (
	"adventofcode2023/days/shared"
	"fmt"
)

type Sequence struct {
	Values []int
}

func (s *Sequence) Length() int {
	return len(s.Values)
}

func (s *Sequence) StepSizes() [][]int {
	stepSizes := [][]int{}
	stepSizes = append(stepSizes, s.Values)
	lineIndex := 1

	for {
		line := []int{}
		foundAllZeroStepSizes := true

		for i := 1; i < len(stepSizes[lineIndex-1]); i++ {
			stepSize := stepSizes[lineIndex-1][i] - stepSizes[lineIndex-1][i-1]

			line = append(line, stepSize)

			if stepSize != 0 {
				foundAllZeroStepSizes = false
			}
		}

		if foundAllZeroStepSizes {
			break
		}

		stepSizes = append(stepSizes, line)

		lineIndex++
	}

	for i := len(stepSizes) - 1; i >= 0; i-- {
		row := stepSizes[i]

		incrementSize := 0

		if i < len(stepSizes)-1 {
			rowUnder := stepSizes[i+1]

			// incrementSize = int(math.Abs(float64(rowUnder[len(rowUnder)-1])))
			incrementSize = rowUnder[len(rowUnder)-1]
		}

		stepSizes[i] = append(row, row[len(row)-1]+incrementSize)
	}

	return stepSizes
}

func (s *Sequence) ReverseStepSizes() [][]int {
	stepSizes := [][]int{}
	stepSizes = append(stepSizes, s.Values)
	lineIndex := 1

	for {
		line := []int{}
		foundAllZeroStepSizes := true

		for i := 1; i < len(stepSizes[lineIndex-1]); i++ {
			stepSize := stepSizes[lineIndex-1][i] - stepSizes[lineIndex-1][i-1]

			line = append(line, stepSize)

			if stepSize != 0 {
				foundAllZeroStepSizes = false
			}
		}

		if foundAllZeroStepSizes {
			break
		}

		stepSizes = append(stepSizes, line)

		lineIndex++
	}

	// For each row the step sizes.
	for i := len(stepSizes) - 1; i >= 0; i-- {
		row := stepSizes[i]

		// Start by assuming it's the bottom row,
		// so increment size is 0.
		incrementSize := 0

		// If it's not the bottom row, then the
		// increment size is first value of the row under.
		if i < len(stepSizes)-1 {
			rowUnder := stepSizes[i+1]

			incrementSize = rowUnder[0]
		}

		stepSizes[i] = append([]int{row[0] - incrementSize}, row...)
	}

	return stepSizes
}

func (s *Sequence) ValuesAsString() string {
	str := ""

	for i, value := range s.Values {
		str += fmt.Sprint(value)

		if i < s.Length()-1 {
			str += " "
		}
	}

	return str
}

func (s *Sequence) StepSizesAsString() []string {
	lines := []string{}

	for _, row := range s.StepSizes() {
		str := ""

		for j, stepSize := range row {
			str += fmt.Sprint(stepSize)

			if j < len(row)-1 {
				str += " "
			}
		}

		lines = append(lines, str)
	}

	return lines
}

func (s *Sequence) DebugStepSizes() {
	for i, line := range s.StepSizesAsString() {
		str := ""

		for j := 0; j < i; j++ {
			str += " "
		}

		println(str + line)
	}
}

func (s *Sequence) FinalValue() int {
	return s.StepSizes()[0][len(s.StepSizes()[0])-1]
}

func (s *Sequence) FinalReversedValue() int {
	return s.ReverseStepSizes()[0][0]
}

func SequenceFromString(line string) Sequence {
	s := Sequence{}

	for _, value := range shared.ExtractIntegersFromString(line) {
		s.Values = append(s.Values, value.Value)
	}

	return s
}
