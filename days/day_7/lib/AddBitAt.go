package lib

func AddBitAt(data uint16, position uint16) uint16 {
	return data + 1<<position
}
