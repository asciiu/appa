package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	_, err := NewPoint(-1, -1, 5, 7)

	assert.Nil(t, err, "err not nil")

	_, err = NewPoint(-1, -2, 5, 7)
	assert.Equal(t, ErrPoint, err, "should be point error")
}
