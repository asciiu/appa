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
