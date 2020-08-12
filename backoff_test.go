package backoff_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xyluet/go-backoff"
)

func Test_Zero(t *testing.T) {
	zeroBackoff := backoff.NewZeroStrategy()
	require.Equal(t, time.Duration(0), zeroBackoff.Duration(0))
	require.Equal(t, time.Duration(0), zeroBackoff.Duration(1))
}

func Test_Linear(t *testing.T) {
	linearBackoff := backoff.NewLinear(time.Second)
	require.Equal(t, time.Second, linearBackoff.Duration(0))
	require.Equal(t, 2*time.Second, linearBackoff.Duration(1))
	require.Equal(t, 3*time.Second, linearBackoff.Duration(2))
}
