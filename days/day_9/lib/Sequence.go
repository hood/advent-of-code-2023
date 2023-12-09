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

func (s *Sequence) StepSizes() []int {
	stepSizes := []int{}

	for i := 1; i < s.Length(); i++ {
		stepSizes = append(stepSizes, s.values[i]-s.values[i-1])
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

func (s *Sequence) StepSizesAsString() string {
	str := ""

	for i, value := range s.StepSizes() {
		str += fmt.Sprint(value)

		if i < len(s.StepSizes())-1 {
			str += " "
		}
	}

	return str
}

func SequenceFromString(line string) Sequence {
	s := Sequence{}

	for _, value := range shared.ExtractIntegersFromString(line) {
		s.values = append(s.values, value.Value)
	}

	return s
}
