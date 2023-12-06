package lib

type Map struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (m *Map) FindMapping(value int) int {
	if value < m.SourceRangeStart {
		return -1
	}

	if value >= m.SourceRangeStart+m.RangeLength {
		return -1
	}

	return m.DestinationRangeStart + (value - m.SourceRangeStart)
}
