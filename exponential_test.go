package backoff_test

import (
	"testing"
	"time"

	"github.com/xyluet/go-backoff"
)

func Test_Exponential(t *testing.T) {
	exp := backoff.NewExponential()
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
			dur := exp.Backoff(tc.retries)
			if tc.expected != dur {
				t.Fatalf("backoff expect: %s, got: %s", tc.expected, dur)
			}
		})
	}
}
