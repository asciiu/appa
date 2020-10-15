package lib

import "github.com/pkg/errors"

var ErrPointNotOnCurve = errors.New("point is not on curve")
var ErrPointMismatch = errors.New("point is not on curve")

// Point on elliptic curve
type Point struct {
	x *int
	y *int
	a int
	b int
}

// NewPoint will return a new point
func NewPoint(x, y *int, a, b int) (*Point, error) {
	point := &Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}

	if x == nil && y == nil {
		return point, nil
	}

	cx := *x
	cy := *y

	if cy*cy != cx*cx*cx+a*cx+b {
		return nil, ErrPointNotOnCurve
	}

	return point, nil
}

// Add will add one finite field to another
func (self Point) Add(other Point) (*Point, error) {
	if self.a != other.a || self.b != other.b {
		return nil, ErrPointMismatch
	}

	if self.x == nil {
		return &other, nil
	}
	if other.x == nil {
		return &self, nil
	}
}
