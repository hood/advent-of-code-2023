
import (
	"fmt"
	"math"
)

// Polynomial represents a polynomial
type Polynomial struct {
	// Coeffs are stored so that Coeffs[len(Coeffs)-1] is the constant
	Coeffs []float64
}

// Evaluate p using Horner's algorithm
func (p Polynomial) Evaluate(x float64) (result float64) {
	for _, coeff := range p.Coeffs {
		result = result*x + coeff
	}
	return result
}

// Derivative returns the derivative of the polynomial
func (p Polynomial) Derivative() (d Polynomial) {
	l := len(p.Coeffs)
	if l <= 1 {
		return d
	}
	d.Coeffs = make([]float64, l-1)
	for i, coeff := range p.Coeffs[:l-1] {
		d.Coeffs[i] = float64(l-i-1) * coeff
	}
	return d
}

// SolveNewton finds a zero using Newton's algorithm, starting at firstGuess until precision is reached
func (p Polynomial) SolveNewton(firstGuess, precision float64) float64 {
	d := p.Derivative()

	newtonDelta := func(x float64) float64 {
		return p.Evaluate(x) / d.Evaluate(x)
	}
	guess := firstGuess
	h := newtonDelta(guess)
	for math.Abs(h) >= precision {
		guess -= h
		h = newtonDelta(guess)
	}
	return guess - h
}

func main() {
	poly := Polynomial{Coeffs: []float64{1, -2, 1, -4}}
	fmt.Println(poly.SolveNewton(float64(2), float64(0.01))) // 2.3146...
}
