package backoff

import (
	"time"
)

// Strategy represents the backoff strategy.
type Strategy interface {
	Backoff(retries int) time.Duration
}

type zero struct{}

// NewZero returns a new zero backoff which always returns 0 duration backoff.
func NewZero() Strategy { return &zero{} }

func (z *zero) Backoff(_ int) time.Duration { return 0 }
