package service

import (
	"errors"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
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

type LogMiddleware struct {
	Logger log.Logger
	Next   StringService
}

func (mw LogMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(s)
	return
}

func (mw LogMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
