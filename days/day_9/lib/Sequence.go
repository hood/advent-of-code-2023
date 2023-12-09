package lib

import (
	"adventofcode2023/days/shared"
	"fmt"
)

type Sequence struct {
	values []int
}

func (s *Sequence) Length() int {
	return len(s.values)
}

func (s *Sequence) StepSizes() [][]int {
	stepSizes := [][]int{}
	stepSizes = append(stepSizes, s.values)
	lineIndex := 1

	for {
		line := []int{}
		sumOfStepSizes := 0

		for i := 1; i < len(stepSizes[lineIndex-1]); i++ {
			stepSize := stepSizes[lineIndex-1][i] - stepSizes[lineIndex-1][i-1]

			line = append(line, stepSize)

			sumOfStepSizes += stepSize
		}

		if sumOfStepSizes == 0 {
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

			incrementSize = rowUnder[len(rowUnder)-1]
		}

		stepSizes[i] = append(row, row[len(row)-1]+incrementSize)
	}

	return stepSizes
}

func (s *Sequence) ValuesAsString() string {
	str := ""

	for i, value := range s.values {
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

func SequenceFromString(line string) Sequence {
	s := Sequence{}

	for _, value := range shared.ExtractIntegersFromString(line) {
		s.values = append(s.values, value.Value)
	}

	return s
}
