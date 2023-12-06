package lib

type Map struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (m *Map) FindMapping(value int) int {
	if value < m.SourceRangeStart {
		return value
	}

	if value >= m.SourceRangeStart+m.RangeLength {
		return value
	}

	return m.DestinationRangeStart + (value - m.SourceRangeStart)
}
