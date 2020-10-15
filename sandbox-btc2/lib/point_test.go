package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	x := -1
	y := -1

	_, err := NewPoint(&x, &y, 5, 7)

	assert.Nil(t, err, "err not nil")

	x2 := -1
	y2 := -2

	_, err = NewPoint(&x2, &y2, 5, 7)
	assert.Equal(t, ErrPointNotOnCurve, err, "should be point error")
}
