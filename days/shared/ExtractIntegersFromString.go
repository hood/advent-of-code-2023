package shared

import "strconv"

func ExtractIntegersFromString(line string) []IntegerWithCoordinates {
	extractor := NewIntegersExtractor()

	for column, char := range line {
		if extractor.ShouldCapture(char) {
			extractor.Capture(char, column)

			if column == len(line)-1 {
				extractor.EndCapture(column)
			}

			continue
		}

		// If was not capturing, continue.
		if !extractor.IsCapturing() {
			continue
		}

		// If it was capturing, but the sequence is over, wrap it up.
		extractor.EndCapture(column)
	}

	return extractor.Values()
}

type IntegersExtractor struct {
	temporaryAccumulator string
	values               []IntegerWithCoordinates
	current              IntegerWithCoordinates
}

func NewIntegersExtractor() *IntegersExtractor {
	return &IntegersExtractor{
		temporaryAccumulator: "",
		values:               []IntegerWithCoordinates{},
		current: IntegerWithCoordinates{
			Value:       -1,
			Row:         -1,
			StartColumn: -1,
			EndColumn:   -1,
		},
	}
}

func (extractor *IntegersExtractor) Values() []IntegerWithCoordinates {
	return extractor.values
}

func (extractor *IntegersExtractor) Capture(char rune, column int) {
	if extractor.current.StartColumn == -1 {
		extractor.startCapture(column)
	}

	extractor.temporaryAccumulator += string(char)
}

func (extractor *IntegersExtractor) startCapture(column int) {
	extractor.current.StartColumn = column
}

func (extractor *IntegersExtractor) IsCapturing() bool {
	return extractor.current.StartColumn != -1
}

func (extractor *IntegersExtractor) EndCapture(column int) {
	// If was capturing, but it's not a number, wrap it up.
	integer, error := strconv.Atoi(extractor.temporaryAccumulator)
	if error != nil {
		panic(error)
	}

	extractor.current.EndColumn = column - 1
	extractor.current.Value = integer

	extractor.values = append(extractor.values, extractor.current)

	extractor.ResetCapture()
}

func (extractor *IntegersExtractor) ShouldCapture(char rune) bool {
	return char >= '0' && char <= '9'
}

func (extractor *IntegersExtractor) ResetCapture() {
	extractor.temporaryAccumulator = ""

	extractor.current = IntegerWithCoordinates{
		Value:       -1,
		StartColumn: -1,
		EndColumn:   -1,
	}
}

func (extractor *IntegersExtractor) Result() []IntegerWithCoordinates {
	return extractor.values
}
