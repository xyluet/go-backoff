package backoff

import "time"

type constant struct {
	duration time.Duration
}

func NewConstant(duration time.Duration) Strategy {
	s := constant{
		duration: duration,
	}
	return &s
}

func (s *constant) Backoff(_ int) time.Duration { return s.duration }
