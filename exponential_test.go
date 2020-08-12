package backoff_test

import (
	"testing"
	"time"

	"github.com/xyluet/go-backoff"
)

func Test_Exponential(t *testing.T) {
	t.Run("default factor (2)", func(t *testing.T) {
		exponentialBackoff := backoff.NewExponentialStrategy(time.Second)
		for name, tc := range map[string]struct {
			retries  int
			expected time.Duration
		}{
			"0": {0, time.Second},
			"1": {1, 2 * time.Second},
			"2": {2, 4 * time.Second},
			"3": {3, 8 * time.Second},
			"4": {4, 16 * time.Second},
			"8": {8, 256 * time.Second},
		} {
			t.Run(name, func(t *testing.T) {
				dur := exponentialBackoff.Duration(tc.retries)
				if tc.expected != dur {
					t.Fatalf("duration expect: %s, got: %s", tc.expected, dur)
				}
			})
		}
	})
}
