package backoff_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xyluet/go-backoff"
)

func Test_Linear(t *testing.T) {
	s := backoff.NewLinear(time.Second)
	require.Equal(t, time.Second, s.Backoff(0))
	require.Equal(t, time.Second, s.Backoff(0))
}
