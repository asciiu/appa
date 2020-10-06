package lib

import (
	"fmt"

	"github.com/pkg/errors"
)

// FieldElement is within a finite field
type FieldElement struct {
	Num   uint
	Prime uint
}

// NewFieldElement creates a new finite field element
func NewFieldElement(num, prime uint) FieldElement {
	if num >= prime || num < 0 {
		panic(errors.Errorf("num %d not in field range 0 to %d", num, prime-1))
	}

	return FieldElement{
		Num:   num,
		Prime: prime,
	}
}

func (fe FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.Prime, fe.Num)
}
