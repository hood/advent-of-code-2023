package lib

func AddBitAt(data uint64, position uint8) uint64 {
	return data + 1<<position
}
