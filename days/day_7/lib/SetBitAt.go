package lib

func SetBitAt(data uint16, position uint16) uint16 {
	return data | 1<<position
}
