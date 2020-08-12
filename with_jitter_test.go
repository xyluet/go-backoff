package backoff_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/xyluet/go-backoff"
)

func Test_DefaultRandomizer(t *testing.T) {
	// r := backoff.NewDefaultRandomizer()
	// min, max := 5*time.Millisecond, 10*time.Millisecond
	// for i := 0; i < 10; i++ {
	// 	durationBetween(t, r.Randomize(min, max), min, max)
	// }
}

func Test_WithJitter(t *testing.T) {
	var strategy backoff.Strategy
	strategy = backoff.NewZeroStrategy()
	strategy = backoff.WithJitter(strategy)

	duration := strategy.Duration(0)
	durationBetween(t, duration, 0, time.Second)
}

func durationBetween(t *testing.T, duration, min, max time.Duration) {
	require.Truef(t, duration >= min && duration <= max, "%s is not between %s and %s", duration, min, max)
}
