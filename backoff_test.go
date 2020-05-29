package backoff_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xyluet/go-backoff"
)

func Test_Zero(t *testing.T) {
	zero := backoff.NewZero()
	require.Equal(t, time.Duration(0), zero.Backoff(0))
	require.Equal(t, time.Duration(0), zero.Backoff(8))
}
