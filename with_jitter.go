package backoff

import (
	"time"
)

type JitterFunc func() time.Duration

func MinMaxJitter(min, max time.Duration) JitterFunc {
	return func() time.Duration {
		return min + time.Duration(seededRand.Int63n(int64(max-min)+1))
	}
}

type withJitter struct {
	Strategy
	jitterFunc JitterFunc
}

// WithJitter returns a new
func WithJitter(s Strategy) Strategy {
	j := withJitter{
		Strategy:   s,
		jitterFunc: MinMaxJitter(0, time.Second),
	}
	return &j
}

func (j *withJitter) Duration(retries int) time.Duration {
	return j.Strategy.Duration(retries) + j.jitterFunc()
}
