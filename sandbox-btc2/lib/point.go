package lib

import "github.com/pkg/errors"

var ErrPoint = errors.New("point is not on curve")

// Point on elliptic curve
type Point struct {
	x int
	y int
	a int
	b int
}

// NewPoint will return a new point
func NewPoint(x, y, a, b int) (*Point, error) {
	if y*y != x*x*x+a*x+b {
		return nil, ErrPoint
	}

	return &Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}, nil
}
