package backoff

import "time"

type linear struct {
	duration time.Duration
}

func NewLinear(duration time.Duration) Strategy {
	s := linear{
		duration: duration,
	}
	return &s
}

func (s *linear) Backoff(retries int) time.Duration {
	return time.Duration(retries+1) * s.duration
}
