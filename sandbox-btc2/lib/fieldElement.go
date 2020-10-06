package lib

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrPrime = errors.New("primes are not equal")

// FieldElement is within a finite field
type FieldElement struct {
	num   uint
	prime uint
}

// NewFieldElement creates a new finite field element
func NewFieldElement(num, prime uint) FieldElement {
	if num >= prime || num < 0 {
		panic(errors.Errorf("num %d not in field range 0 to %d", num, prime-1))
	}

	return FieldElement{
		num:   num,
		prime: prime,
	}
}

func (fe FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.prime, fe.num)
}

// Add will add one finite field to another
func (fe FieldElement) Add(element FieldElement) (*FieldElement, error) {
	if fe.prime != element.prime {
		return nil, ErrPrime
	}

	r := (fe.num + element.num) % fe.prime
	return &FieldElement{
		num:   r,
		prime: fe.prime,
	}, nil
}
