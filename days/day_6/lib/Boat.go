package lib

type Boat struct {
	Speed int
}

func (b *Boat) SetSpeed(speed int) {
	b.Speed = speed
}

func (b *Boat) DistanceTraveledInTime(time int) int {
	return b.Speed * time
}
