package backoff

import (
	"math"
	"time"
)

type ExponentialOption func(*exponentialStrategy)

type exponentialStrategy struct {
	factor  float64
	initial time.Duration
}

// NewExponentialStrategy (TODO)
func NewExponentialStrategy(initial time.Duration, options ...ExponentialOption) Strategy {
	expo := exponentialStrategy{
		factor:  2,
		initial: initial,
	}
	for _, option := range options {
		option(&expo)
	}
	return &expo
}

func (e *exponentialStrategy) Duration(retries int) time.Duration {
	pow := math.Pow(e.factor, float64(retries))
	dur := pow * float64(e.initial)
	return time.Duration(math.Round(dur))
}
