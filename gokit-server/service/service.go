package service

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type GoKitService struct{}

func (GoKitService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (GoKitService) Count(s string) int {
	return len(s)
}
