package lib

type Compass struct {
	Order []rune
	step  int
}

func NewCompass(order []rune) *Compass {
	return &Compass{
		Order: order,
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
