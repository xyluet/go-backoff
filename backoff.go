package backoff

import (
	"time"
)

// Strategy represents the backoff strategy.
type Strategy interface {
	Duration(retries int) time.Duration
}

type zeroStrategy struct{}

// NewZeroStrategy returns a new zero backoff which always returns 0 duration backoff.
func NewZeroStrategy() Strategy { return &zeroStrategy{} }

func (z *zeroStrategy) Duration(_ int) time.Duration { return 0 }

type constantStrategy struct {
	duration time.Duration
}

func NewConstant(duration time.Duration) Strategy {
	s := constantStrategy{
		duration: duration,
	}
	return &s
}

func (s *constantStrategy) Duration(_ int) time.Duration { return s.duration }

type linear struct {
	duration time.Duration
}

func NewLinear(duration time.Duration) Strategy {
	s := linear{
		duration: duration,
	}
	return &s
}

func (s *linear) Duration(retries int) time.Duration {
	return time.Duration(retries+1) * s.duration
}
