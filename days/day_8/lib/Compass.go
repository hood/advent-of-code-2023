package lib

type Compass struct {
	Order []rune
	step  int
}

func NewCompass() *Compass {
	return &Compass{
		Order: make([]rune, 0),
		step:  0,
	}
}

func (c *Compass) NextDirection() rune {
	direction := c.Order[c.step]

	if c.step == len(c.Order)-1 {
		// TODO: Without a reset this can become a steps counter.
		c.step = 0
	} else {
		c.step++
	}

	return direction
}

func (c *Compass) FromLine(line string) *Compass {
	for _, char := range line {
		c.Order = append(c.Order, char)
	}

	return c
}
