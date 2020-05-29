package backoff

import (
	"math"
	"time"
)

type ExponentialOption func(*exponential)

type exponential struct {
	max    time.Duration
	factor float64
}

// NewExponential (TODO)
func NewExponential(options ...ExponentialOption) Strategy {
	expo := exponential{
		factor: 2,
		max:    0,
	}
	for _, option := range options {
		option(&expo)
	}
	return &expo
}

func (e *exponential) Backoff(retries int) time.Duration {
	dur := time.Duration(math.Pow(e.factor, float64(retries))) * time.Second
	if e.max != 0 && e.max < dur {
		dur = e.max
	}

	return dur
}
