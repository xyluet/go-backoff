package backoff

import (
	"math/rand"
	"time"
)

type Randomizer interface {
	Randomize(min, max time.Duration) time.Duration
}

type DefaultRandomizer struct {
	rand *rand.Rand
}

func NewDefaultRandomizer() *DefaultRandomizer {
	r := DefaultRandomizer{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return &r
}

func (r *DefaultRandomizer) Randomize(min, max time.Duration) time.Duration {
	return min + time.Millisecond*time.Duration(r.rand.Int63n(max.Milliseconds()-min.Milliseconds()+1))
}

type WithJitter struct {
	Strategy
	randomizer Randomizer
	min        time.Duration
	max        time.Duration
}

// NewWithJitter returns a new
func NewWithJitter(s Strategy) *WithJitter {
	j := WithJitter{
		Strategy:   s,
		randomizer: NewDefaultRandomizer(),
		min:        0,
		max:        time.Second,
	}
	return &j
}

func (j *WithJitter) Backoff(retries int) time.Duration {
	dur := j.Strategy.Backoff(retries)
	return dur + j.randomizer.Randomize(j.min, j.max)
}
