package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiniteField(t *testing.T) {
	a := NewFieldElement(7, 13)
	b := NewFieldElement(6, 13)
	c := NewFieldElement(7, 13)

	assert.Equal(t, a, c, "should be equal")
	assert.NotEqual(t, a, b, "should not be equal")
}

func TestFiniteFieldAddSuccess(t *testing.T) {
	a := NewFieldElement(7, 13)
	b := NewFieldElement(10, 13)

	d, err := a.Add(b)
	assert.Equal(t, uint(4), d.num, "should be 0")
	assert.Nil(t, err, "error should be nil")
}

func TestFiniteFieldAddError(t *testing.T) {
	a := NewFieldElement(7, 13)
	b := NewFieldElement(6, 15)

	d, err := a.Add(b)
	assert.Equal(t, ErrPrime, err, "should be prime error")
	assert.Nil(t, d, "value should be nil")
}
