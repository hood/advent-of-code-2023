package lib

type Compass struct {
	Order []rune
	step  int
}

func NewCompass() *Compass {
	return &Compass{
		Order: make([]rune, 0),
	}
}

func (c *Compass) NextDirection() rune {
	if c.step == len(c.Order)-1 {
		c.step = 0
	} else {
		c.step++
	}

	return c.Order[c.step]
}

func (c *Compass) FromLine(line string) *Compass {
	for _, char := range line {
		c.Order = append(c.Order, char)
	}

	return c
}
