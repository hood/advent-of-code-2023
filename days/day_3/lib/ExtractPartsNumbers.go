package lib

import "strconv"

func ExtractPartsNumbers(lines []string) []PartNumber {
	extractor := NewPartExtractor()

	for row, line := range lines {
		extractor.ResetCapture()

		for column, char := range line {
			if extractor.ShouldCapture(char) {
				extractor.Capture(char, column)

				if column == len(line)-1 {
					extractor.EndCapture(row, column)
				}

				continue
			}

			// If was not capturing, continue.
			if !extractor.IsCapturing() {
				continue
			}

			// If it was capturing, but the sequence is over, wrap it up.
			extractor.EndCapture(row, column)
		}
	}

	return extractor.Result()
}

type PartExtractor struct {
	temporaryAccumulator string
	partsNumbers         []PartNumber
	current              PartNumber
}

func NewPartExtractor() *PartExtractor {
	return &PartExtractor{
		temporaryAccumulator: "",
		partsNumbers:         []PartNumber{},
		current: PartNumber{
			Value:       -1,
			Row:         -1,
			StartColumn: -1,
			EndColumn:   -1,
		},
	}
}

func (partExtractor *PartExtractor) Capture(char rune, column int) {
	if partExtractor.current.StartColumn == -1 {
		partExtractor.startCapture(column)
	}

	partExtractor.temporaryAccumulator += string(char)
}

func (partExtractor *PartExtractor) startCapture(column int) {
	partExtractor.current.StartColumn = column
}

func (partExtractor *PartExtractor) IsCapturing() bool {
	return partExtractor.current.StartColumn != -1
}

func (partExtractor *PartExtractor) EndCapture(row int, column int) {
	// If was capturing, but it's not a number, wrap it up.
	integer, error := strconv.Atoi(partExtractor.temporaryAccumulator)
	if error != nil {
		panic(error)
	}

	partExtractor.current.EndColumn = column - 1
	partExtractor.current.Value = integer

	partExtractor.partsNumbers = append(partExtractor.partsNumbers, partExtractor.current)

	partExtractor.current = PartNumber{
		Value:       -1,
		Row:         row,
		StartColumn: -1,
		EndColumn:   -1,
	}

	partExtractor.temporaryAccumulator = ""
}

func (partExtractor *PartExtractor) ShouldCapture(char rune) bool {
	return char >= '0' && char <= '9'
}

func (partExtractor *PartExtractor) ResetCapture() {
	partExtractor.temporaryAccumulator = ""

	partExtractor.current = PartNumber{
		Value:       -1,
		Row:         -1,
		StartColumn: -1,
		EndColumn:   -1,
	}
}

func (partExtractor *PartExtractor) Result() []PartNumber {
	return partExtractor.partsNumbers
}
