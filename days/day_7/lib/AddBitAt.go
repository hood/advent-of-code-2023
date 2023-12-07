package lib

// func AddBitAt(data uint64, position uint8) uint64 {
// 	// return data + 1<<(position*4)

// 	// literally add a bit to the left starting
// 	// from the given position.
// 	// example
// 	// 0000 0001
// 	// after adding a bit at position 2
// 	// becomes
// 		 0000 0011
// }

func AddBitAt(data uint64, value uint8) uint64 {
	position := (value) * 4

	for i := uint8(0); i < 4; i++ {
		bitPos := position + i
		bitmask := uint64(1) << bitPos

		if (data & bitmask) == 0 {
			data |= bitmask
			break
		}
	}

	return data
}
